# Agemplate

**Agemplate is Agenda template generator.**

Agemplate automatically generate agenda file from template.  
We often have **regular** meeting and manage it's agenda such as following directory structure by github.

```
2016
└── 4
    ├── 1.md
    ├── 15.md
    └── 8.md
```

You can easily setup agenda file before meething with **Agemplate**.

# Install

```
$ go get github.com/yutakakinjyo/agemplate
```

# Exec

If You execute command of non parameter at 2016/5/1, then date directory is generated.

```
$ agemplate
```

```
2016
└── 5
    └── 1.md
```

Continue, If you execute following command, new dicretory and agenda are added to agenda tree.

```
$ agemplate 2016/4/1.md
```

```
2016
├── 4
│   └── 1.md
└── 5
    └── 1.md
```

# Template

You should prepare template file. Agenda file contens is copied from template file.  
Also you should set path to Environment variable `AGENDA_TEMPLATE` or `.evn` file.  

`.env`
```
AGENDA_TEMPLATE=path/to/template
```

Alternatively, You can speficy option of template file path such as following.

```
$ agemplate -t path/to/template
```

or

```
$ agemplate --template path/to/template
```
