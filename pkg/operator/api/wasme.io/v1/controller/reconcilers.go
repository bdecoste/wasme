// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	wasme_io_v1 "github.com/solo-io/wasme/pkg/operator/api/wasme.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the FilterDeployment Resource.
// implemented by the user
type FilterDeploymentReconciler interface {
	ReconcileFilterDeployment(obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error)
}

// Reconcile deletion events for the FilterDeployment Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type FilterDeploymentDeletionReconciler interface {
	ReconcileFilterDeploymentDeletion(req reconcile.Request)
}

type FilterDeploymentReconcilerFuncs struct {
	OnReconcileFilterDeployment         func(obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error)
	OnReconcileFilterDeploymentDeletion func(req reconcile.Request)
}

func (f *FilterDeploymentReconcilerFuncs) ReconcileFilterDeployment(obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error) {
	if f.OnReconcileFilterDeployment == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFilterDeployment(obj)
}

func (f *FilterDeploymentReconcilerFuncs) ReconcileFilterDeploymentDeletion(req reconcile.Request) {
	if f.OnReconcileFilterDeploymentDeletion == nil {
		return
	}
	f.OnReconcileFilterDeploymentDeletion(req)
}

// Reconcile and finalize the FilterDeployment Resource
// implemented by the user
type FilterDeploymentFinalizer interface {
	FilterDeploymentReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	FilterDeploymentFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeFilterDeployment(obj *wasme_io_v1.FilterDeployment) error
}

type FilterDeploymentReconcileLoop interface {
	RunFilterDeploymentReconciler(ctx context.Context, rec FilterDeploymentReconciler, predicates ...predicate.Predicate) error
}

type filterDeploymentReconcileLoop struct {
	loop reconcile.Loop
}

func NewFilterDeploymentReconcileLoop(name string, mgr manager.Manager) FilterDeploymentReconcileLoop {
	return &filterDeploymentReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &wasme_io_v1.FilterDeployment{}),
	}
}

func (c *filterDeploymentReconcileLoop) RunFilterDeploymentReconciler(ctx context.Context, reconciler FilterDeploymentReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericFilterDeploymentReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(FilterDeploymentFinalizer); ok {
		reconcilerWrapper = genericFilterDeploymentFinalizer{
			genericFilterDeploymentReconciler: genericReconciler,
			finalizingReconciler:              finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericFilterDeploymentHandler implements a generic reconcile.Reconciler
type genericFilterDeploymentReconciler struct {
	reconciler FilterDeploymentReconciler
}

func (r genericFilterDeploymentReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*wasme_io_v1.FilterDeployment)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FilterDeployment handler received event for %T", object)
	}
	return r.reconciler.ReconcileFilterDeployment(obj)
}

func (r genericFilterDeploymentReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(FilterDeploymentDeletionReconciler); ok {
		deletionReconciler.ReconcileFilterDeploymentDeletion(request)
	}
}

// genericFilterDeploymentFinalizer implements a generic reconcile.FinalizingReconciler
type genericFilterDeploymentFinalizer struct {
	genericFilterDeploymentReconciler
	finalizingReconciler FilterDeploymentFinalizer
}

func (r genericFilterDeploymentFinalizer) FinalizerName() string {
	return r.finalizingReconciler.FilterDeploymentFinalizerName()
}

func (r genericFilterDeploymentFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*wasme_io_v1.FilterDeployment)
	if !ok {
		return errors.Errorf("internal error: FilterDeployment handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeFilterDeployment(obj)
}
