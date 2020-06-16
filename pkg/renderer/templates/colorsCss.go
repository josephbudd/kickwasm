package templates

// ColorsCSS is the rendererprocess/css/colors.css template.
const ColorsCSS = `{{$Dot := .}}
/******************************************************************************

    general

******************************************************************************/

html {
  background: transparent;
}

body {
  background: transparent;
}

/******************************************************************************

    master views

******************************************************************************/

#modalInformationMasterView
{
  background: hsla(44, 100%, 19%, 0.8);
  border-color:black;
}

#modalInformationMasterView-center
{
  background-color:white;
  border-color:rgb(88, 65, 1);
}

/******************************************************************************

    panels

******************************************************************************/

div.{{.UnderTabBar}}  
{
  border-color: white;
  background: transparent;
  color:white;
}

/******************************************************************************

    tabs

******************************************************************************/

div.{{.TabBar}} > button
{
  color:black;
  border-color: black;
  background-color:#a09665;
  box-shadow: 2px 0px black;
}
div.{{.TabBar}} > button.selected-tab 
{
  color:black;
  border-color: white;
  background-color:white;
  box-shadow: 2px 0px lightgrey;
}
div.{{.TabBar}} > button.under-tab 
{
  color:black;
  background-color:#e9be02;
  border-color: black;
  box-shadow: 0px;
}

/******************************************************************************

    tabs master view

******************************************************************************/

/******************************************************************************

    home

******************************************************************************/

#{{.IDHomePad}}
{
  border-color:black;
}

/******************************************************************************

    slider

******************************************************************************/

#{{.IDSliderCollection}} > .{{.SliderPanel}} > .{{.SliderPanelPad}}
{
  color:white;
}

/******************************************************************************

  color levels

******************************************************************************/

/* color level 0 if for the home button pad */

#mainMasterView > h1.{{.PanelHeading}}
{
  color: black;
}

button.{{.ClassBackColorLevelPrefix}}0
{
  background-color:black;
  color:white;
  border-color:white;
}
button.{{.ClassBackColorLevelPrefix}}0:hover
{
  color:coral;
  border-color:coral;
}

div.{{.ClassPadColorLevelPrefix}}0
{
  background-color: black;
}

{{range $i, $homeName := .HomeNames}}{{$mod5 := call $Dot.Mod5 $i}}
{{if eq $mod5 0}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color:hsl(50, 23%, 27%);
  color:white;
  padding:5px;
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(50, 23%, 27%);
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(50, 23%, 27%);
  color:hsl(40, 37%, 57%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color:hsl(40, 37%, 57%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(50, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(40, 37%, 57%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(40, 37%, 37%);
  color:white;
  border-color:white;
}
{{else if eq $mod5 1}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(70, 23%, 27%);
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color: hsl(70, 23%, 27%);
  color:white;
  padding:5px;
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(70, 23%, 27%);
  color:hsl(60, 37%, 57%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color:hsl(60, 37%, 57%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(70, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(60, 37%, 57%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(60, 37%, 37%);
  color:white;
  border-color:white;
}

{{else if eq $mod5 2}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(150, 23%, 27%);
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color: hsl(150, 23%, 27%);
  color:white;
  padding:5px;
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(150, 23%, 27%);
  color:hsl(140, 37%, 57%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color:hsl(140, 37%, 57%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(150, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(140, 37%, 57%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(140, 37%, 37%);
  color:white;
  border-color:white;
}

{{else if eq $mod5 3}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(200, 23%, 27%);
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color: hsl(200, 23%, 27%);
  color:white;
  padding:5px;
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(200, 23%, 27%);
  color:hsl(190, 37%, 57%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color:hsl(190, 37%, 57%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(200, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(190, 37%, 57%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(190, 37%, 37%);
  color:white;
  border-color:white;
}

{{else if eq $mod5 4}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(250, 23%, 27%);
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color: hsl(250, 23%, 27%);
  color:white;
  padding:5px;
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(250, 23%, 27%);
  color:hsl(240, 37%, 67%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color: hsl(240, 37%, 67%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(250, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(240, 37%, 67%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(240, 37%, 37%);
  color:white;
  border-color:white;
}

{{else if eq $mod5 5}}
/* color level {{$homeName}} */

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > h2.{{$Dot.ClassPanelHeadingLevelPrefix}}{{$homeName}}
{
  color: hsl(350, 23%, 27%);
}

#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$homeName}}
{
  background-color: hsl(350, 23%, 27%);
  color:white;
  padding:5px;
}

button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}
{
  background-color:hsl(350, 23%, 27%);
  color:hsl(10, 37%, 57%);
  border-color:white;
}
button.{{$Dot.ClassBackColorLevelPrefix}}{{$homeName}}:hover
{
  color:white;
  border-color:hsl(10, 37%, 57%);
}

div.{{$Dot.ClassPadColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(350, 23%, 27%);
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}
{
  background-color: hsl(10, 37%, 57%);
  color:black;
  border-color:black;
}

button.{{$Dot.ClassPadButtonColorLevelPrefix}}{{$homeName}}:hover
{
  background-color: hsl(10, 37%, 37%);
  color:white;
  border-color:white;
}
{{end}}{{end}}

/******************************************************************************

  user content

******************************************************************************/

.{{.UserContent}} a {color:hsla(210, 25%, 80%, 1);}
.{{.UserContent}} a:hover{color:hsla(0, 25%, 80%, 1);}
.{{.UserContent}} a:visited{color:hsla(110, 25%, 80%, 1);}

.{{.UserContent}} button
{
  border-color:black;
  color:black;
  background: #c5b549;
}
.{{.UserContent}} button:hover
{
  background: #554e20;
  color:white;
}

.{{.UserContent}},
.{{.UserContent}} th,
.{{.UserContent}} td
{
  color: white;
}

/******************************************************************************

  modal user content

******************************************************************************/


.{{.ModalUserContent}} a {color:black;}
.{{.ModalUserContent}} a:hover{color:black;}
.{{.ModalUserContent}} a:visited{color:black;}

.{{.ModalUserContent}} button
{
  border-color:black;
  color:black;
  background: white;
}
.{{.ModalUserContent}} button:hover
{
  background: white;
  color:black;
}

.{{.ModalUserContent}},
.{{.ModalUserContent}} th,
.{{.ModalUserContent}} td
{
  background: white;
  color:black;
}
`
