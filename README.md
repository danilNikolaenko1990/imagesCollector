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
- remove dependence on libexif
- tests
- GUI
- windows version