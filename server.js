const fs = require('fs/promises')
const fs2 = require('fs')

require('http').createServer(async (req, res) => {
  if (req.url === '/') {
    const file = await fs.readFile("index.html")
    res.write(file)
    res.end('')
    // fs.createReadStream("index.html").pipe(res)
  }

  if (req.url === '/get') {
    const range = req.headers["range"];
    if (!range) {
      throw new Error("where is range header, my dear?")
    }

    const stats = await fs.stat("fil-paul-talk.mov")
    const start = Number(range.replace(/\D/g, ''));
    const step = 1024 * 1000;
    const end = Math.min(start + step, stats.size - 1)

    const headers = {
      "Content-Range": `bytes ${start}-${end}/${stats.size}`,
      "Accept-Ranges": "bytes",
      "Content-Length": end - start + 1,
      "Content-Type": "video/mp4",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "*",
      "Access-Control-Allow-Methods": "*"
    };

    res.writeHead(206, headers);
    fs2.createReadStream("fil-paul-talk.mov", { start, end }).pipe(res)
    return;
  }

  res.end('')

}).listen(3000, () => { console.log('up on 3000') })