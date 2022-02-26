# Learning Rust Essentials
Watch [this short video](https://www.youtube.com/watch?v=5C_HPTJg5ek) to get the rough grasp of Rust lang
> Rust is a multi-paradigm programming language focused on performance and safety, especially safe concurrency. It is syntactically similar to C++ but provides memory safety without using garbage collection. Rust programming language was developed by Mozilla with the aim of creating a better tool for developing their browser Mozilla Firefox. 

> The language appeared to be so effective, that many programmers are now opting to use it for software development instead of using C++. Rust is syntactically similar to C++, but it provides increased speed and better memory safety.

## Installation
> The first step is to install Rust. We’ll download Rust through rustup, a command line tool for managing Rust versions and associated tools. 

For Mac OS, Linux user, run bellow command. 
```shell
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

For Windows user, follow below link.
> It looks like you’re running Windows. To start using Rust, download the installer, then run the program and follow the onscreen instructions. You may need to install the Visual Studio C++ Build tools when prompted to do so. If you are not on Windows see "Other Installation Methods".

- [install Rust](https://www.rust-lang.org/tools/install)

### Toolchain management with rustup
> Rust is installed and managed by the rustup tool. Rust has a 6-week rapid release process and supports a great number of platforms, so there are many builds of Rust available at any time. rustup manages these builds in a consistent way on every platform that Rust supports, enabling installation of Rust from the beta and nightly release channels as well as support for additional cross-compilation targets.

### Configuring the PATH environment variable
> In the Rust development environment, all tools are installed to the ~/.cargo/bin directory, and this is where you will find the Rust toolchain, including rustc, cargo, and rustup.

> Accordingly, it is customary for Rust developers to include this directory in their PATH environment variable. During installation rustup will attempt to configure the PATH. Because of differences between platforms, command shells, and bugs in rustup, the modifications to PATH may not take effect until the console is restarted, or the user is logged out, or it may not succeed at all.

> If, after installation, running rustc --version in the console fails, this is the most likely reason.

### Updating and uninstalling
To update Rust, 
```shell
$rustup update
```

to uninstall, 
```shell 
$rustup self uninstall
```

## Cargo
> Cargo is Rust’s build system and package manager. Most Rustaceans use this tool to manage their Rust projects because Cargo handles a lot of tasks for you, such as building your code, downloading the libraries your code depends on, and building those libraries. (We call the libraries that your code needs dependencies.)

> Because the vast majority of Rust projects use Cargo, the rest of this book assumes that you’re using Cargo too. Cargo comes installed with Rust if you used the official installers discussed in the “Installation” section. If you installed Rust through some other means, check whether Cargo is installed by entering the following into your terminal:

### First project with cargo 
Run below command to create your first Rust project. 

```shell
$cargo new (project-name) --bin
```

And then write your first Rust function. 

```rs
fn main() {
    println!("hello rust");
}
```

And run below command at your project directory to run the funciton. 

```shell
$cargo run
```

## Reference 
- [dcode : Rust programming tutorial](https://youtube.com/playlist?list=PLVvjrrRCBy2JSHf9tGxGKJ-bYAN_uDCUL)
- [Fireship : Rust in 100 Seconds](https://www.youtube.com/watch?v=5C_HPTJg5ek)