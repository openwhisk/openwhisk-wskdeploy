// Licensed to the Apache Software Foundation (ASF) under one or more contributor
// license agreements; and to You under the Apache License, Version 2.0.

function main(params) {
  console.log("params: " + JSON.stringify(params, null, 4));
  return {payload:  params};
}
