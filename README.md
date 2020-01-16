# gostego

A simple steganography tool written in go, using LSB methodology.

## Usage

### Hide a message

In order to hide a message, pick an image and create a file with your text. Please note that only ASCII characters are supported.

Run the following command to hide the image:

```bash
gostego hide --image myImage.png --msg-file message.txt
```

This will create an image named `myImage_steg.png` in your current folder.

### Show a message

In order to show a message simply run this:

```bash
gostego show --image myImage_steg.png
```

This will return the message. However, since we did not specify how long the included message is, it will also display a bunch of other printable characters. The displayed characters are ASCII characters with character codes between 31 and 123.
If you know how many characters the hidden message has, you can specify that with the `--length` flag:

```bash
gostego show --image myImage_steg.png --length 200
```

This will print out the first 200 printable characters found in the image.
