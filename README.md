This program walk through all the folders you configured and picks images. Then it copies them into target folder 
organized by folders named like 2021/April/IMG-04_10_2018_12_59_15-Apple_Iphone-0a1016.jpg 

How to use:
# OSX
brew install libexif

# Debian
sudo apt-get install -y libexif-dev

Install:
- git clone git@github.com:danilNikolaenko1990/imagesCollector.git
- cp config.yml.dist config.yml
- modify config
- go build
- ./imagesCollector

todo:
- yaml config +
- command (plain folder result, by year, by year and months) +
- console progress bar
- remove libexif dependency
- tests
- GUI
- windows version