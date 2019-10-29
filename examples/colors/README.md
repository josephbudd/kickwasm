# colors

## Designating wide panels in the kickwams.yaml file

Below is the part if this framework's kickwasm.yaml file located at **./kickwasm/yaml/kickwasm.yaml**.

### hvscroll: true

Notice that the panel named **Action1Level1MarkupPanel** has the field **hvscroll** set to **true**. That means that the panel should have not just a default vertical scroll like all panels do but also a horizontal scroll.

Notice also that the panel's markup has a div with the class **"wide-text"** which wraps the rest all of the content.

```yaml
title : Example of the Different Action Colors

importPath : github.com/josephbudd/kickwasm/examples/colors

buttons :
  - name : Action1Button
    label : Action 1 Colors
    heading : Default Action 1 Colors.
    cc : Action 1 Level 1
    panels :
      - name : Action1Level1ButtonPanel
        buttons :
          - name : Action1Level1ContentButton
            label : Click here for level 1 content.
            heading : Markup In Action 1 colors.
            cc : A Markup Panel
            panels :
              - name : Action1Level1MarkupPanel
                hvscroll: true
                note : |
                  This is the only content.
                  Brought to you in the first service color.
                markup : |
                  <div class="wide-text">
                    <h3>A view of Action 1 content.</h3>
                    <h4>A simple and easy to understand interface.
                    ...</h4>
                    <h5>This panel is 1000px wide.</h5>
                    <p>
                    This is an example of content that you would provide.
                    It has the same shape...
                    blah...
                    blah...
                    blah...
                    </p>
                  </div>
```

## Styling wide panels in the css file

### min-width

There fore, the class **wide-text** is styled with the **min-width** set. It is the first setting in ./site/css/Usercontent.css. I added it there becuase Usercontent.css is an editable file. I know that becuase it's name begins with a capital letter.

``` css


/******************************************************************************

  user content
  user markup

******************************************************************************/

.wide-text
{
  min-width: 1000px;
}

...
...
...
```