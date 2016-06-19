# Agemplate

[![Build Status](https://travis-ci.org/yutakakinjyo/agemplate.svg?branch=travis)](https://travis-ci.org/yutakakinjyo/agemplate)

**Agemplate** is **Agenda template generator.**

Agemplate automatically generate agenda file from template.  
We often have **regular** meeting and manage it's agenda such as following directory structure.

```
2016
└── 4
    ├── 1.md
    ├── 15.md
    └── 8.md
```

You can easily setup agenda file before meething with **Agemplate**.

## Getting Started

get agemplate
```
$ go get github.com/yutakakinjyo/agemplate
```

make agendas dicretory to manage agenda files.
```
$ mkdir agendas
$ cd agendas
```

create template file

`template.md`
```
# Agenda

## Member

## Subjects
```

Let's setup agenda.
```
$ agemplate
```

You will get agenda file of meeting. Next, please write the contents.

```
$ tree .
```

```
2016
└── 5
    └── 1.md
```

---

# Install

```
$ go get github.com/yutakakinjyo/agemplate
```

# Exec

If you execute command of non parameter at 2016/5/1, Then date directory is generated.

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
Template file name must be `template.md`. Because **Agemplate** search file named `template.md`.

Alternatively, You can speficy option of template file path such as following.

```
$ agemplate -t path/to/template.md
```

or

```
$ agemplate --template path/to/template.md
```
