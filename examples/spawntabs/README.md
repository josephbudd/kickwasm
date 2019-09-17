## Intro

spawntabs demonstrates how to spawn and unspawn a spawned tab.

A spawned tab is a tab in a tab bar that does not exist until it is spawned and can be removed by unspawning it. When a spawned tab unspawns it destroys itself and all of it's markup panels.

## Summary

1. The markup panel at renderer/panels/TabsButton/TabsButtonTabBarPanel/FirstTab/CreatePanel initiates the spawning of a new renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/. The CreatePanel's controler at renderer/panels/TabsButton/TabsButtonTabBarPanel/FirstTab/CreatePanel/Controller.go
   1. has a button in it's template panel's markup and handles the button's onclick event with the controler's func handleClick.
   1. func handleClick spawns a new renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/ with a call to **secondtab.Spawn(tabLabel, panelHeading, nil)**.
1. SecondTab's one and only markup panel's controler at renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/HelloWorldTemplatePanel/Controller.go
   1. has a button in it's template panel's markup and handles the button's onclick event with the controler's func handleClick.
   1. func handleClick unspawns the the panel's tab.
