---
Title: Sed is So Dead for Regular Expressions
Subtitle: Perl *Destroys* Sed
Query: true
---

If you are trying to get stuff done. Perl destroys `sed` for most
things. Sed cannot do even basic substitution (which people tend to use
`awk` for as well).

```sh
perl -pe 's/foo/bar/'
sed 's/foo/bar/'
perl -pe 's/([fb])ar/\1ar/'
sed 💩
```

## See Also

* [Perl Haters Need to Shut Up and Learn It](/advice/dont/perlhate/)
