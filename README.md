<h1 align="center">
    <br>
    <img src="https://i.imgur.com/iBSb0Fh.png" alt="Baguette#" width="540">
    <br>
    Baguette#
    <br>
</h1>
<h4 align="center">The latest evolution of <a href="https://github.com/hugolgst/ikea-sharp">Ikea#</a></h4>

## Presentation
The *Baguette Sharp* is a fork of Ikea#, the idea is simple, take Ikea# and replace all the Ikea furniture by names of our great 
french desert gastronomy! And you have a *beautiful* not to mention *yummy* language to learn the french gastronomy

## Documentation
You can find examples of the language in the examples folder, here is a program which take 2 integers and print its sum
```baguette sharp
CROISSANT CHOUQUETTE CANELÉ CHOUQUETTE ÉCLAIR CHOUQUETTE CLAFOUTIS ÉCLAIR CHOUQUETTE CLAFOUTIS CLAFOUTIS CLAFOUTIS BAGUETTE
```

### Methods
* Generic
    * CHOUQUETTE : `(`
    * CLAFOUTIS : `)`
    * BAGUETTE : `;`
    * PARISBREST : `"`
* Math
    * CANELÉ [...Integers] : Return the sum of the given integers `CANELÉ CHOUQUETTE 10 12 02 CLAFOUTIS BAGUETTE`
    * PAINAURAISIN [Int Int] : Return the subtraction of the given integers `PAINAURAISIN CHOUQUETTE 10 8 CLAFOUTIS BAGUETTE`
    * STHONORÉ [...Integers] : Return the multiplication of the given integers `STHONORÉ CHOUQUETTE 10 10 10 CLAFOUTIS BAGUETTE`
    * CHOCOLATINE [Int Int] : Return the division of the two given integers `CHOCOLATINE CHOUQUETTE 10 2 CLAFOUTIS BAGUETTE`
    * BRETZEL [Int Int] : Return a random integers between two bounds `BRETZEL CHOUQUETTE 10 8 CLAFOUTIS BAGUETTE`
    * KOUGNAMANN [Int Int] : Return a^b `KOUGNAMANN CHOUQUETTE 10 8 CLAFOUTIS BAGUETTE`
    * BAGUETTEVIÉNOISE [Int Int] : Return the log base b of a `BAGUETTEVIÉNOISE CHOUQUETTE 10 8 CLAFOUTIS BAGUETTE`
    * FINANCIER [Int] : Return the nth fibonacci number `FINANCIER CHOUQUETTE 10 CLAFOUTIS BAGUETTE`
    * PROFITEROLES [Int] : Return the sqrt of the given integer `PROFITEROLES CHOUQUETTE 100 CLAFOUTIS BAGUETTE`
    * OPERA [Int] : Return the opposite of the given integer `OPERA CHOUQUETTE 10 CLAFOUTIS BAGUETTE`
* IO
    * CROISSANT [String] : Println the given parameter `CROISSANT CHOUQUETTE PARISBREST lol PARISBREST CLAFOUTIS BAGUETTE`
    * PAINAUCHOCOLAT [String String] : Printf `PAINAUCHOCOLAT CHOUQUETTE PARISBREST Hello %s! PARISBRST PARISBREST World! PARISBREST CLAFOUTIS BAGUETTE`
    * ÉCLAIR [] : Read the standard input
* Cache
    * QUATREQUART [String a String b] : Save the value b in the variable named a `QUATREQUART CHOUQUETTE PARISBREST a PARISBREST PARISBREST hello PARISBREST CLAFOUTIS BAGUETTE`
    * MADELEINE [String] : Return the value of the variable named a `MADELEINE CHOUQUETTE PARISBREST a PARISBREST CLAFOUTIS BAGUETTE`