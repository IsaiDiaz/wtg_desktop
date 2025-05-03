export namespace category {
	
	export class Category {
	    id: number;
	    name: string;
	    payment: number;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.payment = source["payment"];
	    }
	}

}

export namespace device {
	
	export class Device {
	    id: number;
	    uuid: string;
	    employee_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.uuid = source["uuid"];
	        this.employee_id = source["employee_id"];
	    }
	}

}

export namespace employee {
	
	export class Employee {
	    ID: number;
	    Name: string;
	    CI: string;
	    // Go type: time
	    BirthDate: any;
	    // Go type: time
	    StartDate: any;
	    PhotoURL: string;
	    Auth: number;
	    CategoryID: number;
	    Email: string;
	
	    static createFrom(source: any = {}) {
	        return new Employee(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.CI = source["CI"];
	        this.BirthDate = this.convertValues(source["BirthDate"], null);
	        this.StartDate = this.convertValues(source["StartDate"], null);
	        this.PhotoURL = source["PhotoURL"];
	        this.Auth = source["Auth"];
	        this.CategoryID = source["CategoryID"];
	        this.Email = source["Email"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

