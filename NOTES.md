## What is the plan:

- We are going to send the file
- We are going to hash the file which we will use as a key
- We will create an interface function that will work with the key and help use in creating sub folders (maybe in pairs of two)
- We store the actual encrypted data there

#### Why do we want to use Key/Value pairs?

Hadoop (the framework on which the project is based on) implements the MapReduce paradigm, as mentioned by [Google in their paper on MapReduce]. The output type of the map should match the input type of the reduce

```
(K1,V1) -> Map -> (K2,V2)
(K2, V2) -> Reduce -> (K3,V3)
```

According to the Google's paper:

All and all, The KV approach worked best for Hadoop and we're trying to follow in their steps ;)
