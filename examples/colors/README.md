# What is there to learn from the colors example

## Horizontally scrolled markup panels

**examples/colors/.kickwasm/yaml/kickwasm.yaml** is the kickwasm.yaml file that was used to generate the colors source code. In the kickwasm.yaml file the markup panel **Action1Level1MarkupPanel** has

1. **hvscroll: true** which turns on horizontal scrolling for that panel.
1. **markup** where a **div.wide-text** wraps the rest of the markup.

Below is the relative part of **kickwasm.yaml**.

```yaml

            panels :
              - name : Action1Level1MarkupPanel
                hvscroll: true
                note : |
                  This is the only content.
                  Brought to you in the first service color.
                markup : |
                  <div class="wide-text">
                    <h3>A view of Action 1 content.</h3>
                    <h4>A simple and easy to understand interface.</h4>
                    <h5>This panel is 1000px wide.</h5>
                    <p>
                    This is an example of content that you would provide.
                    It has the same shape and the same background color as this level's button pad.
                    </p>
                    <p>
                    You can see this level's button pad by clicking the tall back button at the left.


```

**examples/colors/site/mycss/Usercontent.css** contains my added style for wide panels. The first panel in each section of this application is wide. Below is that part of **Usercontent.css**.

```css

/******************************************************************************

  user content
  user markup

******************************************************************************/
.wide-text
{
  min-width: 1000px;
}

```

**examples/colors/site/templates/Action1Button/Action1Level1ButtonPanel/Action1Level1ContentButton/Action1Level1MarkupPanel.tmpl** is the template for the first markup panel in the **Action1** section of the application. It's not the first panel in the **Action1** section because there are other button panels. It is just the first markup panel in the **Action1** section.

**Action1Level1MarkupPanel.tmpl** contains the **div.wide-text** which gets styled with a min-width of 1000px as it wraps the other content. The relative part of **Action1Level1MarkupPanel.tmpl** is shown below.

```html

<div class="wide-text">
  <h3>A view of Action 1 content.</h3>
  <h4>A simple and easy to understand interface.</h4>
  <h5>This panel is 1000px wide.</h5>
  <p>
  This is an example of content that you would provide.
  It has the same shape and the same background color as this level's button pad.
  </p>
  <p>
  You can see this level's button pad by clicking the tall back button at the left.

```