# trainee-tail

## About

**trainee-tail** is a from scratch developed program which outputs the last 10 lines of a file you specify. 
Seen like this it's an imitation of the original tail command line command.

## Parameters

| Parameter    | Usage | 
| ------------- |-------------|
| **-n**      | Define how many lines should be printed.

## Usage Examples

In the following you'll see a quick example of how to use the tail command. Here is what our text file looks like:

in.txt:

                one
                two
                three
                four
                five
                six
                seven
                eight
                nine
                ten
                eleven
                twelve
                thirteen
                fourteen
                
To print the last 10 lines of in.txt run `./tail.linux-amd64 < in.txt`

The output is:

                five
                six
                seven
                eight
                nine
                ten
                eleven
                twelve
                thirteen
                fourteen
                
If you want to choose more than one file run `./tail.linux-amd64 in.txt in2.txt in3.txt ...`

To print e.g the last 5 lines use the -n parameter `./tail.linux-amd64 -n 5 in.txt`
