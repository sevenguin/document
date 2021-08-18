module school

go 1.16

// replace sample.com/math => ../math
// 这个在module这一层即可
replace sample.com/learn => ../learn

// require sample.com/math v0.0.0-00010101000000-000000000000

require sample.com/learn v0.0.0-00010101000000-000000000000
