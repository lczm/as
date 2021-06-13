# As

This is a toy programming language mostly built for fun. The repository contains an interpreter to run 'as' code. Performance itself is not very fast as this is a tree-walking interpreter.

## Build
```bash
git clone https://github.com/lczm/as
cd as
go build
```

## Usage
```bash
./as {location_of_file}
```

## Language Details

### Variables
```javascript
var a = 10;
var b = 100;
var c = "hello";
```

### Operations
All your standard `+`, `-`, `*`, `/`, `%` operators
```javascript
var a = 10 + 10;
var b = 10 - 10;
var c = 10 * 10;
var d = 10 / 10;
var e = 10 % 10;
```
Augmented operator assignments, `+=`, `-=`, `*=`, `/=`, `%=` operators
```javascript
var a = 10;
a += 10;
a -= 5;
a *= 2;
a /= 3;
a %= 10;
```
Increment / Decrements
```javascript
var a = 10;
var b = 10;

a++;
b--;
```

### Control Flow & Comparison Operators & Logical Operators
Comparison operators include `<`, `>`, `<=`, `>=`, `==`, `!=`

Logical operators include `&&`, `||`
```javascript
if (10 < 5) {
    print("10 < 5")
} else {
    print("10 > 5")
}
```

### Loops
While loops
```javascript
var a = 0;
while (a < 10) {
    a++;
}
```

For loops
```javascript
for (var i = 0; i < 10; i++) {

}
```

### Functions
```javascript
function fib(n) {
    if (n <= 1) {
        return n;
    }
    return fib(n - 2) + fib(n - 1);
}

var a = fib(5);
```
### Containers
Lists
```javascript
var a = [1, 2, 3];
a = append(a, 4);
a = append(a, 5);
```
HashMaps
```javascript
var a = {0:10, 1: 20};
a[0] = 100;
a[1] = 200;
a[2] = 300;
```

### Structures
```javascript
struct Test {
    var a;
    var b;
    
    init() {
        print("Initialization");
    }
}

var test = Test();
test.a = 10;
print(10);
```

### Builtin Functions
| Functions | Definition                          |
| --------- | ----------------------------------- |
| print()   | Print the out what it the object    |
| len()     | Returns the length of the input     |
| type()    | Returns the type of the input       |
| append()  | Appends an element to the container |

### Examples : Sieve of Eratosthenes
```javascript
function sieve(n) {
    var all = [];
    for (var i = 0; i < n+1; i++) {
        all = append(all, true);
    }

    var p = 2;
    while (p * p <= n) {
        if (all[p] == true) {
            for (var i = p * 2; i < n + 1; i+=p) {
                all[i] = false;
            }
        }
        p += 1;
    }

    all[0] = false;
    all[1] = false;

    for (var i = 0; i < n; i++) {
        if (all[i] == true) {
            print(i);
        }
    }
}

sieve(100);
```

### Notes
Since this is a toy language, just use whatever file extension that works for you. For syntax highlighting, it is more convenient to match the language syntax to Javascript (on whatever editor you use.).