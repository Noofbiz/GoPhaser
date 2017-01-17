# GoPhaser
A way to setup a quick and easy web server using golang to test Phaser.js projects

To use:
  1) Install Go
  2) Clone this Repo to your workspace src folder
  3) Navigate to this folder and run by entering "go run main.go" in your terminal / cmd prompt
  4) Open "localhost:8080" in your browser

You can include your libraries and assets (images, music) in the assets folder.
The most current, stable phaser.min.js library should be added to your assets folder.
The cdn from the example project is version 2.5.0 and likely needs changed for best performance.

The hellophaser example provided is from
  https://github.com/photonstorm/phaser/raw/master/v2/resources/tutorials/01%20Getting%20Started/hellophaser.zip



TODO:
   Get an install script setup so the compiled binaries can be available without needing to install Go
