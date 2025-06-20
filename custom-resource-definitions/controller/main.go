package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

// Website represents our custom resource
type Website struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebsiteSpec   `json:"spec,omitempty"`
	Status WebsiteStatus `json:"status,omitempty"`
}

// DeepCopyObject implements the runtime.Object interface
func (w *Website) DeepCopyObject() runtime.Object {
	if w == nil {
		return nil
	}
	out := new(Website)
	w.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into the given Website
func (w *Website) DeepCopyInto(out *Website) {
	*out = *w
	out.TypeMeta = w.TypeMeta
	out.ObjectMeta = w.ObjectMeta
	out.Spec = w.Spec
	out.Status = w.Status
}

type WebsiteSpec struct {
	Domain   string `json:"domain"`
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
	Port     int32  `json:"port"`
}

type WebsiteStatus struct {
	AvailableReplicas int32  `json:"availableReplicas"`
	Phase             string `json:"phase"`
}

func main() {
	// Get Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fall back to local kubeconfig
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v", err)
		}
	}

	// Create clients
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating kubernetes client: %v", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %v", err)
	}

	// Define the GVR for our custom resource
	gvr := schema.GroupVersionResource{
		Group:    "example.com",
		Version:  "v1",
		Resource: "websites",
	}

	// Create an informer for our custom resource
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return dynamicClient.Resource(gvr).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return dynamicClient.Resource(gvr).Watch(context.TODO(), options)
			},
		},
		&Website{},
		0,
		cache.Indexers{},
	)

	// Add event handlers
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			website := obj.(*Website)
			log.Printf("Website added: %s/%s", website.ObjectMeta.Namespace, website.ObjectMeta.Name)
			// Implement your reconciliation logic here
		},
		UpdateFunc: func(old, new interface{}) {
			_ = old.(*Website) // Ignore old website
			newWebsite := new.(*Website)
			log.Printf("Website updated: %s/%s", newWebsite.ObjectMeta.Namespace, newWebsite.ObjectMeta.Name)
			// Implement your reconciliation logic here
		},
		DeleteFunc: func(obj interface{}) {
			website := obj.(*Website)
			log.Printf("Website deleted: %s/%s", website.ObjectMeta.Namespace, website.ObjectMeta.Name)
			// Implement your cleanup logic here
		},
	})

	// Start the informer
	stopCh := make(chan struct{})
	go informer.Run(stopCh)

	// Wait for the informer to sync
	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		log.Fatal("Failed to sync informer")
	}

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	close(stopCh)
}
