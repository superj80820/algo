async function fetchNeighbors(node: any): Promise<number[]> {
  return new Promise((resolve) => {
    setTimeout(() => {
      switch (node) {
        case 1:
          resolve([2, 3, 4]);
          break;
        case 2:
          resolve([1, 5]);
          break;
        case 3:
          resolve([1, 5]);
          break;
        case 4:
          resolve([1, 6]);
          break;
        case 5:
          resolve([2, 3, 7]);
          break;
        case 6:
          resolve([4, 7]);
          break;
        case 7:
          resolve([5, 6]);
          break;
      }
    }, 1000);
  });
}

// async function searchGraph(start) {
//   queue = [start];
//   visited = new Map();
//   visited[start] = true;
//   while (queue.length > 0) {
//     first = queue.shift();
//     let neighbors = await fetchNeighbors(first);
//     for (let neighbor of neighbors) {
//       if (!visited[neighbor]) {
//         queue.push(neighbor);
//         visited[neighbor] = true;
//       }
//     }
//     console.log(first);
//   }
// }

async function searchGraph(start: any) {
  let visited = new Set([start]);
  let queue = [start];

  while (queue.length) {
    const firstItem = queue.shift();
    console.log(firstItem);
    (await fetchNeighbors(firstItem)).forEach((item) => {
      if (!visited.has(item)) {
        visited.add(item);
        queue.push(item);
      }
    });
  }
}

// async function searchGraphByAsync(start) {
//   queue = [start];
//   visited = new Map();
//   visited[start] = true;

//   while (queue.length > 0) {
//     await Promise.all(
//       queue.map((item) => {
//         console.log(item);
//         return fetchNeighbors(item);
//       })
//     ).then((items) => {
//       queue = [];
//       items.map((neighbors) => {
//         for (let neighbor of neighbors) {
//           if (!visited[neighbor]) {
//             queue.push(neighbor);
//             visited[neighbor] = true;
//           }
//         }
//       });
//     });
//   }
// }

async function searchGraphByAsync(start: any) {
  let visited = new Set([start]);
  let queue = [start];

  while (queue.length) {
    await Promise.all(
      queue.map((item) => {
        console.log(item);
        return fetchNeighbors(item);
      })
    ).then((values) => {
      values.forEach(() => {
        queue.shift();
      });
      values.flat().forEach((item) => {
        if (!visited.has(item)) {
          visited.add(item);
          queue.push(item);
        }
      });
    });
  }
}

async function searchGraphByAsyncV2(node: number, index = 0, batch: Set<number>[] = []) {
  let result = await fetchNeighbors(node);

  const previous = batch[index - 1];

  if (previous) {
    result = result.filter((x) => !previous.has(x));
  }

  (batch[index] ??= new Set()).add(node);

  await Promise.all(result.map((x) => searchGraphByAsyncV2(x, index + 1, batch)));

  return batch;
}

searchGraph(1)
// searchGraphByAsync(1)
// searchGraphByAsyncV2(1).then((batch) => {
//   console.log(batch)
//   return batch.flatMap((x) => Array.from(x))
// })
//   .then(item => {
//     console.log(item)
//   })
//   .catch(console.error);
