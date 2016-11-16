# A Go implementation of MM's Robby program using genetic algorithm

[Link][http://modelai.gettysburg.edu/2016/robby/Chapter9_Complexity.pdf]

Robby is a soda can collection robot that lives in a 2 dimensional 10x10 grid. It's job is to clean up soda cans scatter around the board. It gets rewarded when picking up a can when it's standing on a cell with a soda. And punished when ran into a wall or pick up a can when none is at current site.

Robby is evolved using Genetic algorithm to be smarter an eventually to be very efficient at his job.

My goal is to explore the go language as well as a high level familiarity of how a simple genetic algorithm is implemented. 

To run the trainning program just ```go build && ./train```

Test program is the next thing on the list, I'm interested in visualizing the tests in the [Malmo][https://github.com/Microsoft/malmo] project. Basically bring Robby to the minecraft world!