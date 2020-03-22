# kakuro
Kakuro Solver for learning go

Don't expect anything of quality here.

I am no good at solving kakuros [https://de.wikipedia.org/wiki/Kakuro] but I know a little programming and I want to learn golang. So this I my try at putting it all together.

## Input 
A JSON file stating the dimensions and the fields with sums (constraints) or blocks.

    {
        "w":11,     //width of the field
        "h":11,     //height of the field
        "constraints" : [
            {
                "x" : 0,    //horizontal position
                "y" : 0,    //vertical position
                "h" : -1,   //horizontal sum to the right
                "v" : -1    //vertical sum to the bottom
            },
            ...
        ]
    }

