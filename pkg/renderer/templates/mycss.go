package templates

// MyCSS is the mycss/mycss.css file.
const MyCSS = `{{$Dot := .}}
/******************************************************************************

  print

******************************************************************************/

@media print {
  ::-webkit-scrollbar
  {
    width: 0px;
    height:0px;
  }
  
  #mainMasterView-home-slider
  {
    color: black;
    background: white;
    margin-left: 0px;
  }
  #mainMasterView-home-slider div
  {
    color: black;
    background: white;
  }
  #blackMasterView
  {
    visibility: hidden;
  }
  #mainMasterView-home-slider-back
  {
    visibility: hidden;
  }
  #mainMasterView-home-slider
  {
    margin-left: -50px;
  }
  /* You may or may not want to use these.
  h1,
  h2
  {
    display: none !important;
  } */
}

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

.{{.ModalUserContent}}
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
.{{.ModalUserContent}} td
{
  font-size:20px
}

.{{.UserMarkup}} button,
.{{.ModalUserContent}} button
{
  border-width:1px;
  border-style:solid;
  cursor:pointer;
}
.{{.UserMarkup}} button:hover,
.{{.ModalUserContent}} button:hover
{
  cursor:pointer;
}
.{{.UserMarkup}} button:focus,
.{{.ModalUserContent}} button:focus
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
.{{.ModalUserContent}} h6
{
  font-size:22px;
}

/*
	right margin limits the width and brings it away from the scroll bar.
*/
.{{.UserMarkup}} p,
.{{.ModalUserContent}} p
{
  display:block;
  clear:both;
  width:100%;
  margin-right:20px;
}
`
