# ascii-art-reverse

* Ascii-art-reverse is a program designed to reverse the process of converting text into ASCII art. 
* It  is a tool to convert ASCII art back into readable text. 

## Features

* Convert ASCII art from a file into plain text.
* Handle the --reverse=<fileName> flag to specify the file containing ASCII art.

## Installation

* To use ascii-art-reverse, you need to have Go installed on your system.

* Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/nichotieno/ascii-reverse.git
```

```bash
cd ascii-art-reverse
```

## Usage

### Reversing ASCII Art

* To convert ASCII art from a file back into text, use the --reverse flag followed by the path to the file containing the ASCII art.

```bash
go run . --reverse=<fileName>
```

* Replace <fileName> with the path to your ASCII art file.

#### Example: Reversing ASCII Art

* Suppose you have an ASCII art file named test.txt with the following content:


```bash
 _    _          _   _                __          __                 _       _  
| |  | |        | | | |               \ \        / /                | |     | | 
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | 
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | 
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | 
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| 


```

* Run the following command to convert it back into text:

```bash
go run . --reverse=test.txt
```

##### Output:

```bash
Hello World
```

## Authors

* [Benard Opiyo](https://learn.zone01kisumu.ke/git/beopiyo)

* [Nicholas Otieno](https://learn.zone01kisumu.ke/git/nichotieno)
