const axios = require('axios');

;
(async() => {
  await Promise.all(Array.from(Array(1000).keys())
    .map(() => {
      let data = JSON.stringify({
        "sessionId": "s000000373",
        "refresh": true
      });
      
      let config = {
        method: 'post',
        maxBodyLength: Infinity,
        url: 'https://api.ticketplus.com.tw/captcha/api/v1/generate?_=1698597060000',
        headers: { 
          'authority': 'api.ticketplus.com.tw', 
          'accept': 'application/json, text/plain, */*', 
          'accept-language': 'en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,ja;q=0.6', 
          'authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZldGl4LjE2NDA1MjA4NTAxMzcyMjQiLCJ1c2VyVHlwZSI6Im9vbmUiLCJtYXNrZWRNb2JpbGUiOiIwOTg1KioqNzM4IiwibWFza2VkRW1haWwiOiJzdSoqKioqKioqKkBnbWFpbC5jb20iLCJtYXNrZWRVc2VyTmFtZSI6Iuaelyrlj7MiLCJzZHBJZCI6Ijg1NjExNjczMTkwNDkzIiwidmVyaWZ5RW1haWwiOnRydWUsImVuY3J5cHRlZEFkZHJlc3MiOnsiaXYiOiJhMWI0NGE0OTU1MzEyYjg3ZWUwNGNhZjhlM2Y2M2Y3NyIsImVuY3J5cHRlZERhdGEiOiJkMDdlN2NiODljOGMzYWU1YWNlNjdhNGZhMWI1NjM2YjJhNWIwMjE4MTZjOGE0YTAyNmM2M2IyNDBlYjc0NDNlOTA3ZjE5MzM3Y2ExYzY2NTVkYjkzOWIwMzhlOWViYzI3NGVlZTdkYWUxNzMxZDIzMDdiYzViOTU2ZDFjNjJlMSJ9LCJlbmNyeXB0ZWRVc2VyTmFtZSI6eyJpdiI6IjkxN2JiYzg5MjE2ZGFjZTI1ZjdhMTc2ZmIzNDIzYTNlIiwiZW5jcnlwdGVkRGF0YSI6IjJhZDU3ZGQ4ZDUzN2I4YjliZDNkODhhYWI1MmZjZTc1In0sImRpc2FiaWxpdHkiOm51bGwsImlhdCI6MTY5ODU5NTk2NSwiZXhwIjoxNjk4NTk5NTY1fQ.Aop4zVYBQXzl23u3BoDiTLnCpCnD2obnhzsTVDSMDiI', 
          'content-type': 'application/json', 
          'origin': 'https://ticketplus.com.tw', 
          'referer': 'https://ticketplus.com.tw/', 
          'sec-ch-ua': '"Chromium";v="118", "Google Chrome";v="118", "Not=A?Brand";v="99"', 
          'sec-ch-ua-mobile': '?0', 
          'sec-ch-ua-platform': '"macOS"', 
          'sec-fetch-dest': 'empty', 
          'sec-fetch-mode': 'cors', 
          'sec-fetch-site': 'same-site', 
          'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36'
        },
        data : data
      };
      
      return axios.request(config)
      .then((response) => {
        console.log(JSON.stringify(response.data));
      })
      .catch((error) => {
        console.log(error);
      });
      
    }))
})()


