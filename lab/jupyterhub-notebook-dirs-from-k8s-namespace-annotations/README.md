# Learning Lab: Extend JupyterHub with Menu for Users to Select from One of Several Notebook Directories Pulled from a Kubernetes Namespace Annotation List

JupyterHub is a popular and well-established way to share data science
calculations and models graphically for many people in an organization.
Kubernetes is the leading container orchestration application for such
organizations. 

This lab explores a multi-user environment where the organization has
deployed Kubernetes with network attached storage "scratch" space
(mounted on `/s/<userid>`) for all users matching their user account
ids on a single "login" Linux system. The scratch drive can also be
mounted by anyone in the organization from their personal workstations.
This means some users of the system may never see a shell prompt.

A user's workflow is simply to open an internal web site address, login
with their user id and password, and see a menu of their
`/s/<userid>/notebooks` without being able to see those of other users.

Some users will have several directories to which they have permission
to view notebooks. Therefore, each user should have more than one
possible directory path in which to look for notebooks. Users must be
able to select from a menu (or simply enter the name) of the notebook
directory to set the JupyterHub session to. The list of possible
notebook directories is stored in a Kubernetes namespace that matches
the user id and contains such profile information as annotations or
labels.

For simplification purposes we will not setup NAS but use an
`/s/<userid>` that models the permissions the same way. We'll add a
shared account to simulate additional notebook directory access. This
will greatly simplify the creation of the Kubernetes simulated
deployment.

Also for simplification, the Linux login account credentials will be
used to authenticate even though normally an enterprise-scale web
authentication system (involving Oauth2, Active Directory/LDAP, and/or
Kerberos) would likely be used instead.

## Deliverables

* Git repo containing the following:
  1. Python `get_namespace_annotations(userid)` function that fetches
     the annotations a namespace that matches `userid` *either* from
     within the current Kubernetes context or *out* of cluster looking
     only at `~/.kube/config`
  1. Extended base JupyterHub container suitable for pushing to any
     OCI-compliant container registry (Dockerfile, .dockerignore, etc.)
     1. Add `get_namespace_annotations(userid)`
     1. Add notebook directory menu or text entry field to templates
  1. Kind cluster config that pulls JupyterHub container from local
     container images enabling anyone to do local development with
     quick deployment iterations
  1. Kubernetes Deployment YAML files with comments about where to change things
     from local development, to the staging dev cluster, to production
     1. Main Pod for extended JupyterHub container (most important)
     1. Other services, etc. 
  1. Detailed instructions about how to test, stage, and deploy your
     solution for the customer

## Milestones

These milestones are designed to fit within one week.

* Setup local Kubernetes(Kind)/Docker
* Create env emulation base image
* Add JupyterHub to env base
* Install image as pod and test
* Write and test `get_ns_annotations(ns)`
* Add notebook dir menu to template
* Add user documentation and web text
* Migrate to staging environment
* Conduct systems integration testing
* Create user and support docs and forum
* Complete user-acceptance testing
* Apply user feedback, re-stage, re-test
* Deploy into production
* Report and publish

## Tasks

1. Create a local Kind cluster and point local dev at it

1. Create Linux container (Dockerfile) as base for JupyterHub container
   that will also serve as Python development system:

   1. Create three individual user accounts
   1. Set the passwords for each to `pass{1..3}`
   1. Ensure `/home/<userid>` directories created with proper perms
   1. Create simulated NAS mount as `/s/<userid>` directories
   1. Create three unique notebooks for each in `/s/<userid>/notebooks`
   1. Add a `shared` group
   1. Create a `/s/shared` owned by root but with `shared` group perms
   1. Install Python 3.8+
   1. Build it and run:
      1. Add a bind-mount volume pointing to `cwd` in repo dir

1. Deploy a pod that contains the Linux container
   1. Make sure policy is to update and restart every time the upstream
      "registry" image is updated

1. Deploy namespaces into Kind cluster 
   1. One namespace for each `<userid>` on Linux container
   1. Use `kubectl -f namespaces.yml`

1. Create the Python `get_namespace_annotations(userid)` function using
   the Linux container testing that it works outside and outside of the
   cluster by simply making the changes to files locally in `cwd` and
   building the changes, which will push them into local Docker, which
   will trigger being picked up by Kubernetes and automatically deployed

1. Extend the Linux base container to include JupyterHub being sure to
   mark the separation in the Dockerfile so that only the JupyterHub
   stuff will eventually remain. (The final version of the Dockerfile
   will contain the other stuff commented out and can be uncommented
   when working on the container in development, even though it will
   still have to be setup to be tested from within a cluster with Kind).
   1. Install dependencies
   1. Install JupyterHub
   1. Configure to use `/home/<userid>` logins
