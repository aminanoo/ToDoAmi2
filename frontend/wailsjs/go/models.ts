export namespace models {
	
	export class Task {
	    id: string;
	    title: string;
	    isCompleted: boolean;
	    priority: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.isCompleted = source["isCompleted"];
	        this.priority = source["priority"];
	    }
	}
	export class TaskRequest {
	    id: string;
	    title: string;
	    priority: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.priority = source["priority"];
	    }
	}

}

