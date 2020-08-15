# go-shapes
An example of using delegation to create the appearance of inheritance. An AbstractShape contains the behavior and data common to all shapes. Each shape instance
has its own AbstractShape prototype. Requests to shapes that can be satisfied by the AbstractShape prototype are delegated to that prototype.
