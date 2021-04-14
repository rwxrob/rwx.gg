---
Title: Get WSL Ubuntu on Windows 10
Subtitle: Windows 10 Subsystem for Linux
Query: true
---

The fastest way to get a terminal on Windows 10 is to enable the
*Windows Subsystem for Linux* (commonly called WSL) by using the [built
in installation
command](https://ubuntu.com/blog/new-installation-options-coming-for-ubuntu-wsl).

The rest of this contains the earlier instructions on how to do the same
thing. You may not need them if you do the stuff above correctly.

:::co-fyi
All this WSL activation will be a thing of the past when the next version of Microsoft Windows is released with WSL2 and the new terminal included automatically. All this only applies to Windows 10.
:::

Before you do anything else make sure your work is saved and stop everything else you are doing on the computer because this process requires restarting it.

First we need to install Ubuntu from the Microsoft Store. You can skip the prompt to "Sign In", which isn't required.

Then since you are here go ahead and install the new Windows Terminal (Preview). 

When the Ubuntu install finished click on "Launch".

You'll see a terminal pop up containing something like the following text:

```{.out}
Installing, this may take a few minutes...
WslRegisterDistribution failed with error: 0x8007019e
The Windows Subsystem for Linux optional component is not enabled. Please enable it and try again.
See https://aka.ms/wslinstall for details.
Press any key to continue...
```

Don't worry about the error. It just means that WSL isn't activated, which you will do now. Go ahead and open the <https://docs.microsoft.com/en-us/windows/wsl/install-win10> link and follow the instructions, but here they are again for convenience:

Open PowerShell by typing it into the search and watching for it to pop up as an *App*. DO NOT SELECT THE APP! Notice the "Run as Administrator" to the right and select that instead.

:::co-check
Be sure that you open PowerShell with *Run as Administrator* or it won't work.
:::

You should see a blue PowerShell terminal open up.

Paste the following into the blue PowerShell terminal window at the command line:

```powershell
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
```

This will output something like the following to the screen:

```{.out}
PS C:\Windows\system32> Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
Do you want to restart the computer to complete this operation now?
[Y] Yes  [N] No  [?] Help (default is "Y"):
```

You'll have to accept and reboot.

After the reboot type "Ubuntu" into your Windows search bar to start it up.

You'll see a bunch of text scroll on the screen since Ubuntu is being installed. Wait for it install and be patient. There will be long pauses where nothing happens on the screen.

```{.out}
Installing, this may take a few minutes...
```

After the install completes you will be prompted for a username.

```{.out}
Please create a default UNIX user account. The username does not need to match your Windows username.
For more information visit: https://aka.ms/wslusers
Enter new UNIX username:
```

Pick a username that uses nothing but lowercase letters and numbers (so long as it does not start with one).

Then it will prompt you for a password. IT WILL NOT SHOW YOU TYPING THE PASSWORD BUT IT IS WORKING! This password does not need to be particularly long since only you will be using it on this local system.

```{.out}
passwd: password updated successfully
```

That completes the *basic* Ubuntu installation but you're not done yet.

```{.out}
Installation successful!
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.
```

You *still* need to update and upgrade.

:::co-warning
The first think you should do with *any* new Linux installation is update and upgrade because the version of the operating system installation usually lags behind by several months.
:::

First let's get the latest list of system packages for the update and upgrade.

```sh
sudo apt update
```

You will see a lot of stuff scroll on the screen and this can take a while. Be patient.

```{.out}
...
Fetched 16.5 MB in 15s (1080 kB/s)
Reading package lists... Done
Building dependency tree
Reading state information... Done
10 packages can be upgraded. Run 'apt list --upgradable' to see them.
```

Now it's time to upgrade.

:::co-careful
It's easy to confuse the words `update` with `upgrade`. The two are *very* different. The `update` is used regularly to get a fresh list of the software packages available while `upgrade` does a ***full upgrade of your operating system***. Even though it is a good idea to `upgrade` your system frequently to avoid security vulnerabilities it can be a big deal and you might not want to risk it until you are ready. Just be sure you understand the difference.
:::

Enter the following to start the `upgrade` process.

```sh
sudo apt upgrade
```

You'll be show the packages that will be upgraded and prompted to continue.

```{.out}
...
After this operation, 99.3 kB of additional disk space will be used.
Do you want to continue? [Y/n]y
```

This should take quite a while depending on how old the installation package was. Be patient.

Once this finishes are good to go and can start using it or installing other software that you need.

:::co-warning
When you attempt your first copy and paste it might not go as usual. If so, select the text, move to where you want to paste it, then right click to do the pasting. This is regularly an issue for those doing [OverTheWire](https://overthewire.org).
:::
