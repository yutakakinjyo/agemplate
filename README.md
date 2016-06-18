# Agemplate

**Agemplate is Agenda template generator.**

We often have regular meeting and manage it's agenda such as following directory structure by github.

```
2016
└── 4
    ├── 1.md
    ├── 15.md
    └── 8.md
```

Agemplate automatically generate these agenda file from template.

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

you should prepare template file and set path to Environment variable `AGENDA_TEMPLATE`.

or

exectute wich option following

```
$ agemplate -t /path/to/template.md
```
