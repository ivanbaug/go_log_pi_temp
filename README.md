# Log Raspberry pi cpu temp to thingspeak

This is a quick script to obtain and log the raspberry pi cpu temperature on a platform called thingspeak.
I made it in Go to check out the languaje for the first time, although it's not the best for quick scripting in rpi. The binary file ended up being 6.4 MB which could be better but it's done.

## Usage
The binary has to be compiled on the rpi: `go build log_pi_temp.go`.

The resulting binary has the name `log_pi_temp` . For the program to work it needs a supporting config file `config.json` which has an example in the repo, it just needs the api key from your project in thingspeak.

## Crontab
to execute the binary recurrently I used [crontab](https://www.howtogeek.com/101288/how-to-schedule-tasks-on-linux-an-introduction-to-crontab-files/). The resulting task was:

```bash
# m h  dom mon dow   command
*/15 * * * * /home/pi/go/pitemp/log_pi_temp -config=/home/pi/go/pitemp/config.json
```
