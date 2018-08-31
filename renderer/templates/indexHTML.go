package templates

// IndexHTML is the renderer/templates/main.tmpl template.
const IndexHTML = `<!DOCTYPE html>
<html>

{{.Head}}

<body>

{{.TabsMasterView}}

    <div id="modalInformationMasterView" class="{{.Classes.UnSeen}}">
      <div id="modalInformationMasterView-center" class="{{.Classes.UserContent}}">
        <h1 id="modalInformationMasterView-h1" class="{{.Classes.PanelHeading}}"></h1>
        <div id="modalInformationMasterView-message"></div>
        <p>
          <button id="modalInformationMasterView-close">Close</button>
        </p>
      </div>
	</div> <!-- end of #modalInformationMasterView -->

  <div id="closerMasterView" class="{{.Classes.UnSeen}}">
    <div id="closerMasterView-center" class="{{.Classes.UserContent}}">
      <h2>Close this application.</h2>
      <p>
        Are you certain that you want to close this application?
      </p>
      <p>
        <button id="closerMasterView-cancel">No! Do not close this application.</button>
        <button id="closerMasterView-close">Yes, close this application.</button>
      </p>
    </div>
  </div> <!-- end of closer master view -->

</body>
</html>
`
