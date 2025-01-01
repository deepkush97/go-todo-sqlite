# Simple docker build

```
  docker build -f Dockerfile.simple -t todo-simple .
```

# Multistage docker build

```
  docker build -f Dockerfile.multistage  -t todo-multistage .
```

# Optimised multistage docker build

```
  docker build -f Dockerfile.optimised  -t todo-optimised .
```

| Name            | Tag    | Image ID     | Created        | Size          |
| --------------- | ------ | ------------ | -------------- | ------------- |
| todo-simple     | latest | 5f29ed1d268c | 4 minutes ago  | **1.15 GB**   |
| todo-multistage | latest | e62d7936b82c | 38 minutes ago | **859.73 MB** |
| todo-optimised  | latest | 0f93c5afca95 | 31 minutes ago | **23.64 MB**  |
