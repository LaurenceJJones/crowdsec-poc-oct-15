## crowdsec-poc-oct-15
This is just for fun, crowdsec detects via sysmon files created with custom extension (.enc) then kills the PID

### Folders:
//`mal`\\\\ NOT USED ANYMORE Checkout ransomware simulator [here](https://github.com/NextronSystems/ransomware-simulator)</br>
`crowdsec` has custom config files for crowdsec on windows.</br>
`notification-pid` is a custom notification plugin that takes the pid from the Alert object and kills the process.</br>

Few limitations
  - Detects one PID but could be altered to detect a parent PID that is generating a lot of child processes very fast.
  - Because leaky bucket has to one PID
 
