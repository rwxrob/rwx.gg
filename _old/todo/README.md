---
Title: TODO List
Subtitle: RWX Stuff That Needs Work
---

Here are the upcoming planned changes and new content to RWX.GG. This TODO list is maintained instead of an issues system on GitLab only because it is far easier to maintain and regularly access. As priority of TODO items changes they are elevated to the top. Past items are removed entirely or transferred to the [Changes](/changes/) overview.

1. Title capitalization for all FAQ questions.

1. Change services links to image icons

1. Fix line break indicator  in vim-pandoc plugin

1. Copy the YAML knowledge node data model from robs.io blog to `/contrib/guide/meta/` or maybe `/what/knowledge/node/`.

1. Convert *every* node to new data model YAML.

1. Create a `yamlindexlong` that extracts the `Summary` and `Video` link as well as the `Title` and `ID` for each node ID listed. 

1. Add progressive JavaScript that "closes" the `Summary` unless hovered or tapped. When tapped a second time (Summary showing) follow the link.

1. Update to WSL2.

1. Edit values and views

1. Perl pie and other walkthroughs 

1. Highlight and publish videos for week of June 1

1. Include how to setup the Minecraft client along with the server.

1. Add last updated to MANIFEST

1. Add a `changes` page that lists all the nodes by title in reverse order.

1. Walkthrough of copy and paste options with TMUX / Vim / Mouse / Browser.

1. Make an index of all quotes with attribution.

1. Rant about how bad `awk` is in `/stupid`.

1. Remove all `../` parent URLs.

1. Get rid of `magic` in `magic wand` references cuz *some people* think nsfw.

1. `co-rage` -> `co-mad` and check all the others

1. Make a shortcut `_redirect` creator that only asks about conflicting shortcuts and had an ignore list.

1. Change everything to singular.

1. Test on Android device.

1. Review all for consistent use of FAQ styling

1. Write an `audit` script that combines `check` with validation of consistent use of [callouts](/contrib/guide/) and any other style guide violations that might creep in. Have check only run on files that have changed since the last run of `check`.

1. Add *spoiler* mode that will automatically hide all the code rather than give away the commands so that it can be used to review as you go.

1. Check all `sh` code blocks to ensure they shouldn't be `bash` instead. Check `pandoc --list-highlight-languages` to be sure of others.

1. Read global knowledge base data from `/README.md` YAML

## Long Term

1. Vim config to fill in relative local links with fully qualified node
   refs.

1. Think about scenario where nodes are moved and how would affect README.world subscribers if at all without `_redirect`.

1. Need a *full* Git primer in the early stages of the boost.

1. Replace [TLCL](/reviews/books/tlcl/) with *Linux Terminal Mastery*, a knowledge app that can double as a book, PDF, or whatever. The world needs a *true* terminal mastery book that weaves together deep knowledge and skill in using the fastest human-computer interface physically possible. The rather shallow and inexperienced "Keyboard Tricks" chapter in TLCL was really the trigger. It is disastrously bad and uninformed.
