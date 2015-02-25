#  **go_log_rotate**

## adjustment
1. when log error, error output both into files(for production use) and console (for debug use)
2. when log info, info output into console (for debug use only)
3. file containing the log can be rotated, with the help of lumberjake

## with the help of 
1. lumberjake [https://github.com/natefinch/lumberjack]
2. logrus  [https://github.com/Sirupsen/logrus]

I do these adjustment to logrus, making it more friendly to project.	
