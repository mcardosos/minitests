add this to gulpfile
```
var goCommands = [
  'glide install',
  'go test ./go-acceptance-tests/... -v'
];

function goTest() {
  process.env.GOPATH = basePathOrThrow() + '/AutoRest/Generators/Go/Go.Tests';
  return shell.task(goCommands, {
      cwd: './AutoRest/Generators/Go/Go.Tests/src/Tests', verbosity: 3});
}

gulp.task('test:go', goTest());
```