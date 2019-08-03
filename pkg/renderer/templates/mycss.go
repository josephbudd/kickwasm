package templates

// MyCSS is the mycss/mycss.css file.
const MyCSS = `{{$Dot := .}}

/******************************************************************************

  user content
  user markup

******************************************************************************/

.{{.VScroll}}
{
  overflow-x:hidden;
  overflow-y:auto;
}

.{{.HVScroll}}
{
  overflow:auto;
}

.{{.UserContent}}
{
  /* "{{.UserContent}}" is always used with "{{.VScroll}}" or "{{.HVScroll}}" */
  /* It always wraps "{{.UserMarkup}}" */
  padding-right:30px;
}

.{{.UserMarkup}}
{
  /* "{{.UserMarkup}}" only wraps user markup templates */
  letter-spacing: 2px;
  word-spacing: 4px;
}

.{{.ModalUserContent}},
.{{.CloserUserContent}}
{
  overflow-x:hidden;
  overflow-y:auto;
  padding-right:10px;
  letter-spacing: 2px;
  word-spacing: 4px;
}

.{{.UserMarkup}} button,
.{{.UserMarkup}} input,
.{{.UserMarkup}} label,
.{{.UserMarkup}} legend,
.{{.UserMarkup}} select,
.{{.UserMarkup}} textarea,
.{{.UserMarkup}} th,
.{{.UserMarkup}} td,
.{{.ModalUserContent}} button,
.{{.ModalUserContent}} input,
.{{.ModalUserContent}} label,
.{{.ModalUserContent}} legend,
.{{.ModalUserContent}} select,
.{{.ModalUserContent}} textarea,
.{{.ModalUserContent}} th,
.{{.ModalUserContent}} td,
.{{.CloserUserContent}} button,
.{{.CloserUserContent}} input,
.{{.CloserUserContent}} label,
.{{.CloserUserContent}} legend,
.{{.CloserUserContent}} select,
.{{.CloserUserContent}} textarea,
.{{.CloserUserContent}} th,
.{{.CloserUserContent}} td
{
  font-size:20px
}

.{{.UserMarkup}} button,
.{{.ModalUserContent}} button,
.{{.CloserUserContent}} button
{
  border-width:1px;
  border-style:solid;
  cursor:pointer;
}
.{{.UserMarkup}} button:hover,
.{{.ModalUserContent}} button:hover,
.{{.CloserUserContent}} button:hover
{
  cursor:pointer;
}
.{{.UserMarkup}} button:focus,
.{{.ModalUserContent}} button:focus,
.{{.CloserUserContent}} button:focus
{
  outline: none;
}

.{{.UserMarkup}} h3,
.{{.UserMarkup}} h4,
.{{.UserMarkup}} h5,
.{{.UserMarkup}} h6,
.{{.ModalUserContent}} h3,
.{{.ModalUserContent}} h4,
.{{.ModalUserContent}} h5,
.{{.ModalUserContent}} h6,
.{{.CloserUserContent}} h3,
.{{.CloserUserContent}} h4,
.{{.CloserUserContent}} h5,
.{{.CloserUserContent}} h6
{
  font-size:22px;
}

/*
	right margin limits the width and brings it away from the scroll bar.
*/
.{{.UserMarkup}} p,
.{{.ModalUserContent}} p,
.{{.CloserUserContent}} p
{
  display:block;
  clear:both;
  width:100%;
  margin-right:20px;
}
`
