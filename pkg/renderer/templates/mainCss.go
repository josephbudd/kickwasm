package templates

// MainCSS is the main css template.
const MainCSS = `{{$Dot := .}}

/******************************************************************************

    WARNING:

    Back this file up before editing it.

    Colors are in css/colors.css, not this file.

******************************************************************************/

/******************************************************************************

    general

******************************************************************************/

body {
  margin:0;
  padding:10px;
  overflow:hidden;
  font-family: Helvetica, Arial, sans-serif;
  font-size:20px;
}

.{{.UnSeen}} { display:none !important; }
.{{.Seen}} { display:block !important; }

/******************************************************************************

    master views

******************************************************************************/

#tabsMasterView
{
  height:inherit;
  overflow:hidden;
}

#modalInformationMasterView
{
  z-index:2;
  top:10px;
  left:10px;
  position:absolute;
  border-width:1px;
  border-style:solid;
  border-radius: 20px;
  padding: 50px;
}

#modalInformationMasterView-center
{
  position:relative;
  border-width:1px;
  border-style:solid;
  border-radius: 20px;
  padding:10px;
}

#modalInformationMasterView-center > h2
{
  margin-top:0;
  margin-bottom:5px;
  font-size:36px;
}

#modalInformationMasterView-message
{
  overflow-x: hidden;
  overflow-y: auto;
}

#closerMasterView
{
  position: relative;
  z-index:2;
  top:0px;
  left:0px;
  position:absolute;
  border-width:1px;
  border-style:solid;
  border-radius: 20px;
  padding:50px;
}

#closerMasterView-center
{
  position:relative;
  border-width:1px;
  border-style:solid;
  border-radius: 20px;
  padding:10px;
}

#closerMasterView-center > h2
{
  margin-top:0;
  margin-bottom:5px;
  font-size:36px;
}

/******************************************************************************

    panels

******************************************************************************/

div.{{.UnderTabBar}}  
{
  border-width: 1px;
  border-style: solid;
  padding:0px;
  overflow:hidden;
  margin-top:-1px;
}

div.{{.PanelWithHeading}},
div.{{.PanelWithTabBar}}
{
  overflow:hidden;
}

div.{{.TabPanel}}
{
  padding:0px 2px 0px 2px;
}

div.{{.TabPanelGroup}}
{
  overflow-y:auto;
  padding-left:20px;
}

/******************************************************************************

    tabs

******************************************************************************/

div.{{.TabBar}}
{
  align-content: baseline;
  overflow: hidden;
}
div.{{.TabBar}} > button
{
  float: left;
  width: max-intrinsic;
  border-width: 1px;
  border-style: solid;
  border-bottom: none;
  border-radius: 4px 4px 0 0;
  padding:4px 12px;
  margin:9px 0 0 0;
  font-size:16px;
  cursor:pointer;
  border-spacing: 4px;
}
div.{{.TabBar}} > button:focus { outline: none; }
div.{{.TabBar}} > button.selected-tab 
{
  position:relative;
  margin:0 0 -1px 0;
  padding:4px 12px 6px 12px;
  font-size:26px;
  cursor:default;
}

div.{{.UnderTabBar}} .{{.PanelHeading}}
{
	margin-top:5px;
  margin-bottom:5px;
  padding: 10px;
}

/******************************************************************************

    tabs master view

******************************************************************************/

#tabsMasterView > h1.{{.PanelHeading}}
{
  margin-top:0;
  margin-bottom:5px;
  font-size:36px;
}

/******************************************************************************

    home

******************************************************************************/

#{{.IDHomePad}}
{
  overflow-y:auto;
  border-radius: 20px;
  border-width:1px;
  border-style:solid;
}

#{{.IDHomePad}} > button
{
  height:200px;
  width:300px;
  font-size: 200%;
  text-align: center;
  vertical-align: middle;
  margin:40px;
  border-radius: 20px 20px 20px 20px;
  border-width:1px;
  border-style:solid;
  cursor:pointer;
}
#{{.IDHomePad}} > button:focus { outline: none; }

/******************************************************************************

    slider

******************************************************************************/

#{{.IDSlider}}
{
  float:left;
}

#{{.IDSliderBack}}
{
  width:50px;
  float: left;
  margin-right:10px;
  border-radius:50px 0px 0px 50px;
  border-width:1px;
  border-style:solid;
  font-size:200%;
  font-weight:bold;
  cursor:pointer;
}
#{{.IDSliderBack}}:focus { outline: none; }

#{{.IDSlider}}
{
  float: left;
}

#{{.IDSliderCollection}}
{
  float:right;
}

/*
#{{.IDSliderCollection}} > .{{.SliderPanel}}
{
}
*/

#{{.IDSliderCollection}} > .{{.SliderPanel}} > .{{.SliderPanelPad}}
{
  overflow:hidden;
  /* top and bottom padding for scroll bars */
  padding-top:20px;
  padding-bottom:20px;
  padding-left:10px;
  padding-right:10px;
  border-radius: 20px;
}

#{{.IDSliderCollection}} > .{{.SliderPanel}} > .{{.SliderPanelPad}} > .{{.SliderButtonPad}}
{
  overflow-x:hidden;
  overflow-y:auto;
  padding-right:10px;
  letter-spacing: 2px;
  word-spacing: 4px;
}

#{{.IDSliderCollection}} > .{{.SliderPanel}} > .{{.SliderPanelPad}} > .{{.SliderButtonPad}} > button
{
  height:150px;
  width:225px;
  font-size: 150%;
  text-align: center;
  vertical-align: middle;
  margin:40px;
  border-width:1px;
  border-style:solid;
  border-radius: 15px;
}
#{{.IDSliderCollection}} > .{{.SliderPanel}} > .{{.SliderPanelPad}} > .{{.SliderButtonPad}} > button:focus { outline: none; }

#{{.IDSliderCollection}} > .{{.SliderPanel}} > h2.{{.PanelHeading}}
{
  margin-top:0;
  margin-bottom:0px;
  font-size:34px;
}

#{{.IDSliderCollection}} > .{{.SliderPanel}} > div.{{.PanelHeading}}
{
  margin-bottom:10px;
}

/* cookie crumbs */
{{range $i, $serviceName := .ServiceNames}}
#{{$Dot.IDSliderCollection}} > .{{$Dot.SliderPanel}} > div.{{$Dot.PanelHeading}} > h2.{{$Dot.ClassCookieCrumbLevelPrefix}}{{$serviceName}}{{if lt $i $Dot.LastServiceIndex}},{{end}}{{end}}
{
  font-size:24px;
  display:inline;
  margin-right: 4px;
}
`
