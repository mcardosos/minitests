`go test ./... -v`

Add this to gulpfile...

`gulp.task('test:go', shell.task('go test ./... -v', {cwd: './src/generator/AutoRest.Go.Tests', verbosity: 3}));`