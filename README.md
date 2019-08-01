# DO_IT

A tool to increase work productivity by prioritising task, purely writtern in [Go language](http://www.golang.org/) and [Revel Framework](https://revel.github.io/).

### Start the application:

- Initialse the MySql and put the creds in init.sh
- do `source init.sh`
- run `revel run do_it`

### Tech Used
    
    - Revel 
    - Go Lang 
    - HTML templating
    - MySql

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Future Plan

- Integrate Pomodoro, make user accountable by showing current progress.
- Add Friends
- Generate user profile, history.
- Ability to move uncompleted task to current day.
