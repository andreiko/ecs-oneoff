# ecs-oneoff

```
usage: ecs-oneoff --taskdef=TASKDEF [<flags>] [<override>...]

Flags:
  --help               Show context-sensitive help (also try --help-long and --help-man).
  --taskdef=TASKDEF    Family and/or revision of the task definition to run
  --cluster="default"  Cluster on which to run your task
  --count=1            The number of instantiations of each task to place on your cluster

Args:
  [<override>]  name of one or more override files in JSON format
```

## Environment variables

**AWS_ACCESS_KEY_ID**

**AWS_SECRET_ACCESS_KEY**

**AWS_REGION**
