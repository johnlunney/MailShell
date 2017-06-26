extern crate atarashii_imap;
extern crate openssl;

use std::process::Command;
use std::path::Path;
use atarashii_imap::{Connection, Response};
use openssl::ssl::{SslContext, SslStream};
use openssl::ssl::{SslMethod, SslConnectorBuilder};

#[cfg(target_os = "linux")]
fn run_cmd(cmd: String, args: String) -> Option<String> {
    let output = match Command::new(cmd).arg("-c").arg(args).output() {
        Ok(val) => val,
        Err(_) => "",
    };
    String::from_utf8(output.stdout).unwrap()
}

// powershell only
#[cfg(target_os = "windows")]
fn run_cmd(cmd: String) -> Option<String> {
    unimplemented!()
}

fn check_email(username: String, password: String) -> Option<String> {
    let mut conn = match (Connection::open_secure("imap.gmail.com", "gmail_login@gmail.com", "password")){
        Ok(l) => l,
        Err(_) => None
    }
    
}

fn reply_back(response: String) {
    unimplemented!()
}

fn main() {
    loop {
        let cmd = check_email("".to_string(), "".to_string()).unwrap();
        run_cmd("".to_string(), "".to_string());
    }
}
