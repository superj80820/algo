(async() => {
  console.log("foo")
  await new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve()
    }, 1000 * 5)
  })
  console.log("bar")
})()