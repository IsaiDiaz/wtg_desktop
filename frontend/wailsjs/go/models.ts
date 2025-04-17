export namespace employee {
	
	export class Employee {
	    id: number;
	    name: string;
	    ci: string;
	    // Go type: time
	    birth_date: any;
	    // Go type: time
	    start_date: any;
	    photo_url: string;
	    auth: number;
	    category_id: number;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new Employee(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.ci = source["ci"];
	        this.birth_date = this.convertValues(source["birth_date"], null);
	        this.start_date = this.convertValues(source["start_date"], null);
	        this.photo_url = source["photo_url"];
	        this.auth = source["auth"];
	        this.category_id = source["category_id"];
	        this.email = source["email"];
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

