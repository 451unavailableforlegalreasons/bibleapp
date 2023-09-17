class Settings {
    // no setter or getter for settings, just write to it and call writetolocalStorage or loadfromlocalStorage
    constructor() {
        this.settings = {
            hasselectededition: false,
            preferededition: {id: null, name:null, lang:null},
            preferences: {layout: "scroll"}, // or page
            lastvisitdata: {
                edition: null,
                book: 0,
                chapter: 0,
                firstdisplayedverse: 0,
            },
            // ... other fields in browser localstorage
        }
    }


    loadfromlocalStorage() {
        // it's called parseloadingdata because that how i use it in the app
        let pageloadingdata = localStorage.getItem("pageloadingdata")
        console.log(pageloadingdata)
        if (pageloadingdata === null || pageloadingdata === undefined) {
            // init with empty settings in localStorage
            let pageloadstr = JSON.stringify(this.settings)
            console.log("stri:", pageloadstr)
            localStorage.setItem("pageloadingdata", pageloadstr)
        } else {
            let parsedpageloadingdata = JSON.parse(pageloadingdata)
            console.log("parsed:", parsedpageloadingdata)
            this.settings = parsedpageloadingdata 
        }
    }
    writetolocalStorage() {
        let pageloadstr = JSON.stringify(this.settings)
        localStorage.setItem("pageloadingdata", pageloadstr)
    }
}



export { Settings }
