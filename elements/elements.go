// Copyright 2018-19 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package elements

import (
	"github.com/rocketlaunchr/react"

	"github.com/gopherjs/gopherjs/js"
)

var (
	// TRUE is used when a pointer to a bool is required.
	TRUE = true
	// FALSE is used when a pointer to a bool is required.
	FALSE = false
)

// AProps ...
type AProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	//////////

	Href   string `react:"href,omitempty"`
	Target string `react:"target,omitempty"`
}

// A ...
func A(props *AProps, children ...interface{}) *js.Object {
	return react.JSX("a", props, children...)
}

// AbbrProps ...
type AbbrProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Abbr ...
func Abbr(props *AbbrProps, children ...interface{}) *js.Object {
	return react.JSX("abbr", props, children...)
}

// ArticleProps ...
type ArticleProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Article ...
func Article(props *AbbrProps, children ...interface{}) *js.Object {
	return react.JSX("article", props, children...)
}

// AsideProps ...
type AsideProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Aside ...
func Aside(props *AsideProps, children ...interface{}) *js.Object {
	return react.JSX("aside", props, children...)
}

// BProps ...
type BProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// B ...
func B(props *BProps, children ...interface{}) *js.Object {
	return react.JSX("b", props, children...)
}

// BodyProps ...
type BodyProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnAfterPrint   *js.Object `react:"onAfterPrint,omitempty"`
	OnBeforePrint  *js.Object `react:"onBeforePrint,omitempty"`
	OnBeforeUnload *js.Object `react:"onBeforeUnload,omitempty"`
	OnError        *js.Object `react:"onError,omitempty"`
	OnHashChange   *js.Object `react:"onHashChange,omitempty"`
	OnLoad         *js.Object `react:"onLoad,omitempty"`
	OnOffline      *js.Object `react:"onOffline,omitempty"`
	OnOnline       *js.Object `react:"onOnline,omitempty"`
	OnPageShow     *js.Object `react:"onPageShow,omitempty"`
	OnResize       *js.Object `react:"onResize,omitempty"`
	OnUnload       *js.Object `react:"onUnload,omitempty"`
	OnBlur         *js.Object `react:"onBlur,omitempty"`
	OnFocus        *js.Object `react:"onFocus,omitempty"`
	OnChange       *js.Object `react:"onChange,omitempty"`
	OnClick        *js.Object `react:"onClick,omitempty"`
}

// Body ...
func Body(props *BodyProps, children ...interface{}) *js.Object {
	return react.JSX("body", props, children...)
}

// BrProps ...
type BrProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Br ...
func Br(props *BrProps, children ...interface{}) *js.Object {
	return react.JSX("br", props, children...)
}

// ButtonProps ...
type ButtonProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	//////////

	AutoFocus         *bool  `react:"autofocus,omitempty"`
	Disabled          *bool  `react:"disabled,omitempty"`
	Form              string `react:"form,omitempty"`
	FormAction        string `react:"formaction,omitempty"`
	FormEncType       string `react:"formenctype,omitempty"`
	FormMethod        string `react:"formmethod,omitempty"`
	FormNoValidate    *bool  `react:"formnovalidate,omitempty"`
	FormTarget        string `react:"formtarget,omitempty"`
	Name              string `react:"name,omitempty"`
	Type              string `react:"type,omitempty"`
	ValidationMessage string `react:"validationmessage,omitempty"`
	Value             string `react:"value,omitempty"`
	WillValidate      *bool  `react:"willvalidate,omitempty"`
}

// Button ...
func Button(props *ButtonProps, children ...interface{}) *js.Object {
	return react.JSX("button", props, children...)
}

// CaptionProps ...
type CaptionProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Caption ...
func Caption(props *CaptionProps, children ...interface{}) *js.Object {
	return react.JSX("caption", props, children...)
}

// CodeProps ...
type CodeProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Code ...
func Code(props *CodeProps, children ...interface{}) *js.Object {
	return react.JSX("code", props, children...)
}

// DivProps ...
type DivProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Div ...
func Div(props *DivProps, children ...interface{}) *js.Object {
	return react.JSX("div", props, children...)
}

// EmProps ...
type EmProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Em ...
func Em(props *EmProps, children ...interface{}) *js.Object {
	return react.JSX("em", props, children...)
}

// FooterProps ...
type FooterProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Footer ...
func Footer(props *FooterProps, children ...interface{}) *js.Object {
	return react.JSX("footer", props, children...)
}

// FormProps ...
type FormProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur        *js.Object `react:"onBlur,omitempty"`
	OnChange      *js.Object `react:"onChange,omitempty"`
	OnClick       *js.Object `react:"onClick,omitempty"`
	OnContextMenu *js.Object `react:"onContextMenu,omitempty"`
	OnFocus       *js.Object `react:"onFocus,omitempty"`
	OnInput       *js.Object `react:"onInput,omitempty"`
	OnInvalid     *js.Object `react:"onInvalid,omitempty"`
	OnReset       *js.Object `react:"onReset,omitempty"`
	OnSearch      *js.Object `react:"onSearch,omitempty"`
	OnSelect      *js.Object `react:"onSelect,omitempty"`
	OnSubmit      *js.Object `react:"onSubmit,omitempty"`

	AcceptCharset string `react:"acceptCharset,omitempty"`
	Action        string `react:"action,omitempty"`
	Autocomplete  string `react:"autoComplete,omitempty"`
	Encoding      string `react:"encoding,omitempty"`
	Enctype       string `react:"encType,omitempty"`
	Length        *int   `react:"length,omitempty"`
	Method        string `react:"method,omitempty"`
	Name          string `react:"name,omitempty"`
	NoValidate    *bool  `react:"noValidate,omitempty"`
	Target        string `react:"target,omitempty"`
}

// Form ...
func Form(props *FormProps, children ...interface{}) *js.Object {
	return react.JSX("form", props, children...)
}

// H1Props ...
type H1Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H1 ...
func H1(props *H1Props, children ...interface{}) *js.Object {
	return react.JSX("h1", props, children...)
}

// H2Props ...
type H2Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H2 ...
func H2(props *H2Props, children ...interface{}) *js.Object {
	return react.JSX("h2", props, children...)
}

// H3Props ...
type H3Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H3 ...
func H3(props *H3Props, children ...interface{}) *js.Object {
	return react.JSX("h3", props, children...)
}

// H4Props ...
type H4Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H4 ...
func H4(props *H4Props, children ...interface{}) *js.Object {
	return react.JSX("h4", props, children...)
}

// H5Props ...
type H5Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H5 ...
func H5(props *H5Props, children ...interface{}) *js.Object {
	return react.JSX("h5", props, children...)
}

// H6Props ...
type H6Props struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// H6 ...
func H6(props *H6Props, children ...interface{}) *js.Object {
	return react.JSX("h6", props, children...)
}

// HeaderProps ...
type HeaderProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Header ...
func Header(props *HeaderProps, children ...interface{}) *js.Object {
	return react.JSX("header", props, children...)
}

// HrProps ...
type HrProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Hr ...
func Hr(props *HrProps, children ...interface{}) *js.Object {
	return react.JSX("hr", props, children...)
}

// IProps ...
type IProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// I ...
func I(props *IProps, children ...interface{}) *js.Object {
	return react.JSX("i", props, children...)
}

// IFrameProps ...
type IFrameProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	Width    string `react:"width,omitempty"`
	Height   string `react:"height,omitempty"`
	Name     string `react:"name,omitempty"`
	Src      string `react:"src,omitempty"`
	SrcDoc   string `react:"srcDoc,omitempty"`
	Seamless *bool  `react:"seamless,omitempty"`
}

// IFrame ...
func IFrame(props *IFrameProps, children ...interface{}) *js.Object {
	return react.JSX("iframe", props, children...)
}

// ImgProps ...
type ImgProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	Alt           string `react:"alt,omitempty"`
	Complete      *bool  `react:"complete,omitempty"`
	CrossOrigin   string `react:"crossOrigin,omitempty"`
	Height        string `react:"height,omitempty"`
	IsMap         *bool  `react:"isMap,omitempty"`
	NaturalHeight string `react:"naturalHeight,omitempty"`
	NaturalWidth  string `react:"naturalWidth,omitempty"`
	Src           string `react:"src,omitempty"`
	UseMap        string `react:"useMap,omitempty"`
	Width         string `react:"width,omitempty"`
}

// Img ...
func Img(props *ImgProps, children ...interface{}) *js.Object {
	return react.JSX("img", props, children...)
}

// InputProps ...
type InputProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	Accept             string `react:"accept,omitempty"`
	Alt                string `react:"alt,omitempty"`
	Autocomplete       string `react:"autoComplete,omitempty"`
	Autofocus          *bool  `react:"autoFocus,omitempty"`
	Checked            *bool  `react:"checked,omitempty"`
	DefaultChecked     *bool  `react:"defaultChecked,omitempty"`
	DefaultValue       string `react:"defaultValue,omitempty"`
	DirName            string `react:"dirName,omitempty"`
	Disabled           *bool  `react:"disabled,omitempty"`
	FormAction         string `react:"formAction,omitempty"`
	FormEncType        string `react:"formEncType,omitempty"`
	FormMethod         string `react:"formMethod,omitempty"`
	FormNoValidate     *bool  `react:"formNoValidate,omitempty"`
	FormTarget         string `react:"formTarget,omitempty"`
	Height             string `react:"height,omitempty"`
	Indeterminate      *bool  `react:"indeterminate,omitempty"`
	Max                string `react:"max,omitempty"`
	MaxLength          *int   `react:"maxLength,omitempty"`
	Min                string `react:"min,omitempty"`
	Multiple           *bool  `react:"multiple,omitempty"`
	Name               string `react:"name,omitempty"`
	Pattern            string `react:"pattern,omitempty"`
	Placeholder        string `react:"placeholder,omitempty"`
	ReadOnly           *bool  `react:"readOnly,omitempty"`
	Required           *bool  `react:"required,omitempty"`
	SelectionDirection string `react:"selectionDirection,omitempty"`
	SelectionEnd       *int   `react:"selectionEnd,omitempty"`
	SelectionStart     *int   `react:"selectionStart,omitempty"`
	Size               *int   `react:"size,omitempty"`
	Src                string `react:"src,omitempty"`
	Step               string `react:"step,omitempty"`
	Type               string `react:"type,omitempty"`
	ValidationMessage  string `react:"validationMessage,omitempty"`
	Value              string `react:"value,omitempty"`
	Width              string `react:"width,omitempty"`
	WillValidate       *bool  `react:"willValidate,omitempty"`
}

// Input ...
func Input(props *InputProps, children ...interface{}) *js.Object {
	return react.JSX("input", props, children...)
}

// LabelProps ...
type LabelProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	For  string `react:"htmlFor,omitempty"`
	Form string `react:"form,omitempty"`
}

// Label ...
func Label(props *LabelProps, children ...interface{}) *js.Object {
	return react.JSX("label", props, children...)
}

// LiProps ...
type LiProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Li ...
func Li(props *LiProps, children ...interface{}) *js.Object {
	return react.JSX("li", props, children...)
}

// MainProps ...
type MainProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Main ...
func Main(props *MainProps, children ...interface{}) *js.Object {
	return react.JSX("main", props, children...)
}

// NavProps ...
type NavProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Nav ...
func Nav(props *NavProps, children ...interface{}) *js.Object {
	return react.JSX("nav", props, children...)
}

// OptionProps ...
type OptionProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	DefaultSelected *bool  `react:"defaultSelected,omitempty"`
	Disabled        *bool  `react:"disabled,omitempty"`
	Index           *int   `react:"index,omitempty"`
	Label           string `react:"label,omitempty"`
	Selected        *bool  `react:"selected,omitempty"`
	Text            string `react:"text,omitempty"`
	Value           string `react:"value,omitempty"`
}

// Option ...
func Option(props *OptionProps, children ...interface{}) *js.Object {
	return react.JSX("option", props, children...)
}

// PProps ...
type PProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// P ...
func P(props *PProps, children ...interface{}) *js.Object {
	return react.JSX("p", props, children...)
}

// PreProps ...
type PreProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Pre ...
func Pre(props *PreProps, children ...interface{}) *js.Object {
	return react.JSX("pre", props, children...)
}

// SelectProps ...
type SelectProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	Autofocus         *bool  `react:"autofocus,omitempty"`
	Disabled          *bool  `react:"disabled,omitempty"`
	Length            *int   `react:"length,omitempty"`
	Multiple          *bool  `react:"multiple,omitempty"`
	Name              string `react:"name,omitempty"`
	Required          *bool  `react:"required,omitempty"`
	SelectedIndex     *int   `react:"selectedIndex,omitempty"`
	Size              *int   `react:"size,omitempty"`
	Type              string `react:"type,omitempty"`
	ValidationMessage string `react:"validationMessage,omitempty"`
	Value             string `react:"value,omitempty"`
	WillValidate      *bool  `react:"willValidate,omitempty"`
}

// Select ...
func Select(props *SelectProps, children ...interface{}) *js.Object {
	return react.JSX("select", props, children...)
}

// SpanProps ...
type SpanProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Span ...
func Span(props *SpanProps, children ...interface{}) *js.Object {
	return react.JSX("span", props, children...)
}

// SProps ...
type SProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// S ...
func S(props *SProps, children ...interface{}) *js.Object {
	return react.JSX("s", props, children...)
}

// SupProps ...
type SupProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Sup ...
func Sup(props *SupProps, children ...interface{}) *js.Object {
	return react.JSX("sup", props, children...)
}

// TableProps ...
type TableProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Table ...
func Table(props *TableProps, children ...interface{}) *js.Object {
	return react.JSX("table", props, children...)
}

// TBodyProps ...
type TBodyProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// TBody ...
func TBody(props *TBodyProps, children ...interface{}) *js.Object {
	return react.JSX("tbody", props, children...)
}

// TdProps ...
type TdProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Td ...
func Td(props *TdProps, children ...interface{}) *js.Object {
	return react.JSX("td", props, children...)
}

// TextAreaProps ...
type TextAreaProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	AutoComplete       string `react:"autoComplete,omitempty"`
	AutoFocus          *bool  `react:"autoFocus,omitempty"`
	Cols               *int   `react:"cols,omitempty"`
	DefaultValue       string `react:"defaultValue,omitempty"`
	DirName            string `react:"dirName,omitempty"`
	Disabled           *bool  `react:"disabled,omitempty"`
	MaxLength          *int   `react:"maxLength,omitempty"`
	Name               string `react:"name,omitempty"`
	Placeholder        string `react:"placeholder,omitempty"`
	ReadOnly           *bool  `react:"readOnly,omitempty"`
	Required           *bool  `react:"required,omitempty"`
	Rows               *int   `react:"rows,omitempty"`
	SelectionDirection string `react:"selectionDirection,omitempty"`
	SelectionStart     *int   `react:"selectionStart,omitempty"`
	SelectionEnd       *int   `react:"selectionEnd,omitempty"`
	TextLength         *int   `react:"textLength,omitempty"`
	Type               string `react:"type,omitempty"`
	ValidationMessage  string `react:"validationMessage,omitempty"`
	Value              string `react:"value,omitempty"`
	WillValidate       *bool  `react:"willValidate,omitempty"`
	Wrap               string `react:"wrap,omitempty"`
}

// TextArea ...
func TextArea(props *TextAreaProps, children ...interface{}) *js.Object {
	return react.JSX("textarea", props, children...)
}

// TFootProps ...
type TFootProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// TFoot ...
func TFoot(props *TFootProps, children ...interface{}) *js.Object {
	return react.JSX("tfoot", props, children...)
}

// ThProps ...
type ThProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`

	Abbr    string `react:"abbr,omitempty"`
	ColSpan *int   `react:"colSpan,omitempty"`
	Headers string `react:"headers,omitempty"`
	RowSpan *int   `react:"rowSpan,omitempty"`
	Scope   string `react:"scope,omitempty"`
	Sorted  string `react:"sorted,omitempty"`
}

// Th ...
func Th(props *ThProps, children ...interface{}) *js.Object {
	return react.JSX("th", props, children...)
}

// THeadProps ...
type THeadProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// THead ...
func THead(props *THeadProps, children ...interface{}) *js.Object {
	return react.JSX("thead", props, children...)
}

// TrProps ...
type TrProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Tr ...
func Tr(props *TrProps, children ...interface{}) *js.Object {
	return react.JSX("tr", props, children...)
}

// UlProps ...
type UlProps struct {
	AriaSet                 react.Set   `react:"aria-,omitempty"`
	DataSet                 react.Set   `react:"data-,omitempty"`
	DangerouslySetInnerHTML interface{} `react:"dangerouslySetInnerHTML,omitempty"`
	Accesskey               string      `react:"accesskey,omitempty"`
	Class                   string      `react:"className,omitempty"`
	Contenteditable         *bool       `react:"contenteditable,omitempty"`
	Dir                     string      `react:"dir,omitempty"`
	Draggable               *bool       `react:"draggable,omitempty"`
	Dropzone                string      `react:"dropzone,omitempty"`
	Hidden                  *bool       `react:"hidden,omitempty"`
	ID                      string      `react:"id,omitempty"`
	Lang                    string      `react:"lang,omitempty"`
	SpellCheck              *bool       `react:"spellcheck,omitempty"`
	TabIndex                string      `react:"tabindex,omitempty"`
	Title                   string      `react:"title,omitempty"`

	Key   string     `react:"key,omitempty"`
	Ref   *js.Object `react:"ref,omitempty"`
	Role  string     `react:"role,omitempty"`
	Style *Styles    `react:"style,omitempty"`

	// Events
	OnBlur   *js.Object `react:"onBlur,omitempty"`
	OnFocus  *js.Object `react:"onFocus,omitempty"`
	OnChange *js.Object `react:"onChange,omitempty"`
	OnClick  *js.Object `react:"onClick,omitempty"`
}

// Ul ...
func Ul(props *UlProps, children ...interface{}) *js.Object {
	return react.JSX("ul", props, children...)
}
