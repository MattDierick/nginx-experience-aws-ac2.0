// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControl) DeepCopyInto(out *AccessControl) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControl.
func (in *AccessControl) DeepCopy() *AccessControl {
	if in == nil {
		return nil
	}
	out := new(AccessControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(ActionRedirect)
		**out = **in
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(ActionReturn)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(ActionProxy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionProxy) DeepCopyInto(out *ActionProxy) {
	*out = *in
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = new(ProxyRequestHeaders)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseHeaders != nil {
		in, out := &in.ResponseHeaders, &out.ResponseHeaders
		*out = new(ProxyResponseHeaders)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionProxy.
func (in *ActionProxy) DeepCopy() *ActionProxy {
	if in == nil {
		return nil
	}
	out := new(ActionProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionRedirect) DeepCopyInto(out *ActionRedirect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionRedirect.
func (in *ActionRedirect) DeepCopy() *ActionRedirect {
	if in == nil {
		return nil
	}
	out := new(ActionRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionReturn) DeepCopyInto(out *ActionReturn) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionReturn.
func (in *ActionReturn) DeepCopy() *ActionReturn {
	if in == nil {
		return nil
	}
	out := new(ActionReturn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddHeader) DeepCopyInto(out *AddHeader) {
	*out = *in
	out.Header = in.Header
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddHeader.
func (in *AddHeader) DeepCopy() *AddHeader {
	if in == nil {
		return nil
	}
	out := new(AddHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressMTLS) DeepCopyInto(out *EgressMTLS) {
	*out = *in
	if in.VerifyDepth != nil {
		in, out := &in.VerifyDepth, &out.VerifyDepth
		*out = new(int)
		**out = **in
	}
	if in.SessionReuse != nil {
		in, out := &in.SessionReuse, &out.SessionReuse
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressMTLS.
func (in *EgressMTLS) DeepCopy() *EgressMTLS {
	if in == nil {
		return nil
	}
	out := new(EgressMTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPage) DeepCopyInto(out *ErrorPage) {
	*out = *in
	if in.Codes != nil {
		in, out := &in.Codes, &out.Codes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(ErrorPageReturn)
		(*in).DeepCopyInto(*out)
	}
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(ErrorPageRedirect)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPage.
func (in *ErrorPage) DeepCopy() *ErrorPage {
	if in == nil {
		return nil
	}
	out := new(ErrorPage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPageRedirect) DeepCopyInto(out *ErrorPageRedirect) {
	*out = *in
	out.ActionRedirect = in.ActionRedirect
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPageRedirect.
func (in *ErrorPageRedirect) DeepCopy() *ErrorPageRedirect {
	if in == nil {
		return nil
	}
	out := new(ErrorPageRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPageReturn) DeepCopyInto(out *ErrorPageReturn) {
	*out = *in
	out.ActionReturn = in.ActionReturn
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPageReturn.
func (in *ErrorPageReturn) DeepCopy() *ErrorPageReturn {
	if in == nil {
		return nil
	}
	out := new(ErrorPageReturn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEndpoint) DeepCopyInto(out *ExternalEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEndpoint.
func (in *ExternalEndpoint) DeepCopy() *ExternalEndpoint {
	if in == nil {
		return nil
	}
	out := new(ExternalEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Header) DeepCopyInto(out *Header) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(UpstreamTLS)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMTLS) DeepCopyInto(out *IngressMTLS) {
	*out = *in
	if in.VerifyDepth != nil {
		in, out := &in.VerifyDepth, &out.VerifyDepth
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMTLS.
func (in *IngressMTLS) DeepCopy() *IngressMTLS {
	if in == nil {
		return nil
	}
	out := new(IngressMTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTAuth) DeepCopyInto(out *JWTAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTAuth.
func (in *JWTAuth) DeepCopy() *JWTAuth {
	if in == nil {
		return nil
	}
	out := new(JWTAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Match) DeepCopyInto(out *Match) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	if in.Splits != nil {
		in, out := &in.Splits, &out.Splits
		*out = make([]Split, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDC) DeepCopyInto(out *OIDC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDC.
func (in *OIDC) DeepCopy() *OIDC {
	if in == nil {
		return nil
	}
	out := new(OIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyReference) DeepCopyInto(out *PolicyReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyReference.
func (in *PolicyReference) DeepCopy() *PolicyReference {
	if in == nil {
		return nil
	}
	out := new(PolicyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.AccessControl != nil {
		in, out := &in.AccessControl, &out.AccessControl
		*out = new(AccessControl)
		(*in).DeepCopyInto(*out)
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(RateLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.JWTAuth != nil {
		in, out := &in.JWTAuth, &out.JWTAuth
		*out = new(JWTAuth)
		**out = **in
	}
	if in.IngressMTLS != nil {
		in, out := &in.IngressMTLS, &out.IngressMTLS
		*out = new(IngressMTLS)
		(*in).DeepCopyInto(*out)
	}
	if in.EgressMTLS != nil {
		in, out := &in.EgressMTLS, &out.EgressMTLS
		*out = new(EgressMTLS)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(OIDC)
		**out = **in
	}
	if in.WAF != nil {
		in, out := &in.WAF, &out.WAF
		*out = new(WAF)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyRequestHeaders) DeepCopyInto(out *ProxyRequestHeaders) {
	*out = *in
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = new(bool)
		**out = **in
	}
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyRequestHeaders.
func (in *ProxyRequestHeaders) DeepCopy() *ProxyRequestHeaders {
	if in == nil {
		return nil
	}
	out := new(ProxyRequestHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyResponseHeaders) DeepCopyInto(out *ProxyResponseHeaders) {
	*out = *in
	if in.Hide != nil {
		in, out := &in.Hide, &out.Hide
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]AddHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyResponseHeaders.
func (in *ProxyResponseHeaders) DeepCopy() *ProxyResponseHeaders {
	if in == nil {
		return nil
	}
	out := new(ProxyResponseHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	*out = *in
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(int)
		**out = **in
	}
	if in.NoDelay != nil {
		in, out := &in.NoDelay, &out.NoDelay
		*out = new(bool)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.RejectCode != nil {
		in, out := &in.RejectCode, &out.RejectCode
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]PolicyReference, len(*in))
		copy(*out, *in)
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	if in.Splits != nil {
		in, out := &in.Splits, &out.Splits
		*out = make([]Split, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]Match, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ErrorPages != nil {
		in, out := &in.ErrorPages, &out.ErrorPages
		*out = make([]ErrorPage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityLog) DeepCopyInto(out *SecurityLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityLog.
func (in *SecurityLog) DeepCopy() *SecurityLog {
	if in == nil {
		return nil
	}
	out := new(SecurityLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionCookie) DeepCopyInto(out *SessionCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionCookie.
func (in *SessionCookie) DeepCopy() *SessionCookie {
	if in == nil {
		return nil
	}
	out := new(SessionCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Split) DeepCopyInto(out *Split) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Split.
func (in *Split) DeepCopy() *Split {
	if in == nil {
		return nil
	}
	out := new(Split)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(TLSRedirect)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSRedirect) DeepCopyInto(out *TLSRedirect) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSRedirect.
func (in *TLSRedirect) DeepCopy() *TLSRedirect {
	if in == nil {
		return nil
	}
	out := new(TLSRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upstream) DeepCopyInto(out *Upstream) {
	*out = *in
	if in.Subselector != nil {
		in, out := &in.Subselector, &out.Subselector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaxFails != nil {
		in, out := &in.MaxFails, &out.MaxFails
		*out = new(int)
		**out = **in
	}
	if in.MaxConns != nil {
		in, out := &in.MaxConns, &out.MaxConns
		*out = new(int)
		**out = **in
	}
	if in.Keepalive != nil {
		in, out := &in.Keepalive, &out.Keepalive
		*out = new(int)
		**out = **in
	}
	if in.ProxyBuffering != nil {
		in, out := &in.ProxyBuffering, &out.ProxyBuffering
		*out = new(bool)
		**out = **in
	}
	if in.ProxyBuffers != nil {
		in, out := &in.ProxyBuffers, &out.ProxyBuffers
		*out = new(UpstreamBuffers)
		**out = **in
	}
	out.TLS = in.TLS
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Queue != nil {
		in, out := &in.Queue, &out.Queue
		*out = new(UpstreamQueue)
		**out = **in
	}
	if in.SessionCookie != nil {
		in, out := &in.SessionCookie, &out.SessionCookie
		*out = new(SessionCookie)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upstream.
func (in *Upstream) DeepCopy() *Upstream {
	if in == nil {
		return nil
	}
	out := new(Upstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamBuffers) DeepCopyInto(out *UpstreamBuffers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamBuffers.
func (in *UpstreamBuffers) DeepCopy() *UpstreamBuffers {
	if in == nil {
		return nil
	}
	out := new(UpstreamBuffers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamQueue) DeepCopyInto(out *UpstreamQueue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamQueue.
func (in *UpstreamQueue) DeepCopy() *UpstreamQueue {
	if in == nil {
		return nil
	}
	out := new(UpstreamQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTLS) DeepCopyInto(out *UpstreamTLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTLS.
func (in *UpstreamTLS) DeepCopy() *UpstreamTLS {
	if in == nil {
		return nil
	}
	out := new(UpstreamTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServer) DeepCopyInto(out *VirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServer.
func (in *VirtualServer) DeepCopy() *VirtualServer {
	if in == nil {
		return nil
	}
	out := new(VirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerList) DeepCopyInto(out *VirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerList.
func (in *VirtualServerList) DeepCopy() *VirtualServerList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRoute) DeepCopyInto(out *VirtualServerRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRoute.
func (in *VirtualServerRoute) DeepCopy() *VirtualServerRoute {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteList) DeepCopyInto(out *VirtualServerRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServerRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteList.
func (in *VirtualServerRouteList) DeepCopy() *VirtualServerRouteList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteSpec) DeepCopyInto(out *VirtualServerRouteSpec) {
	*out = *in
	if in.Upstreams != nil {
		in, out := &in.Upstreams, &out.Upstreams
		*out = make([]Upstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subroutes != nil {
		in, out := &in.Subroutes, &out.Subroutes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteSpec.
func (in *VirtualServerRouteSpec) DeepCopy() *VirtualServerRouteSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteStatus) DeepCopyInto(out *VirtualServerRouteStatus) {
	*out = *in
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]ExternalEndpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteStatus.
func (in *VirtualServerRouteStatus) DeepCopy() *VirtualServerRouteStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerSpec) DeepCopyInto(out *VirtualServerSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]PolicyReference, len(*in))
		copy(*out, *in)
	}
	if in.Upstreams != nil {
		in, out := &in.Upstreams, &out.Upstreams
		*out = make([]Upstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerSpec.
func (in *VirtualServerSpec) DeepCopy() *VirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStatus) DeepCopyInto(out *VirtualServerStatus) {
	*out = *in
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]ExternalEndpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStatus.
func (in *VirtualServerStatus) DeepCopy() *VirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WAF) DeepCopyInto(out *WAF) {
	*out = *in
	if in.SecurityLog != nil {
		in, out := &in.SecurityLog, &out.SecurityLog
		*out = new(SecurityLog)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WAF.
func (in *WAF) DeepCopy() *WAF {
	if in == nil {
		return nil
	}
	out := new(WAF)
	in.DeepCopyInto(out)
	return out
}