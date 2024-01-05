# tinygo-keyboard-examples

This repository is an example of a project using sago35/tinygo-keyboard.

You can write to the Wio Terminal using the following command.


```
$ tinygo flash --target wioterminal --size short .
```

## how to start your keyboard project

You can use the "gonew" command to copy the project template.
Alternatively, you can also clone the GitHub project or download it as a zip file.

```
$ gonew github.com/sago35/tinygo-keyboard-examples example.com/your/keyboard
gonew: initialized example.com/your/keyboard in .\keyboard

$ cd ./keyboard

$ tinygo flash --target wioterminal --size short .
```
