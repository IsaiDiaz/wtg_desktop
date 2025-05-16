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
	    Phone: string;
	
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
	        this.Phone = source["Phone"];
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

export namespace project {
	
	export class Project {
	    ID: number;
	    Name: string;
	    // Go type: time
	    InitialDate: any;
	    // Go type: time
	    FinalDate?: any;
	    IsCurrent: boolean;
	    Status: boolean;
	    Description: string;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.InitialDate = this.convertValues(source["InitialDate"], null);
	        this.FinalDate = this.convertValues(source["FinalDate"], null);
	        this.IsCurrent = source["IsCurrent"];
	        this.Status = source["Status"];
	        this.Description = source["Description"];
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

export namespace projectemployee {
	
	export class ProjectEmployee {
	    ID: number;
	    ProjectID: number;
	    EmployeeID: number;
	    CurrentCategoryID: number;
	    Status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProjectEmployee(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ProjectID = source["ProjectID"];
	        this.EmployeeID = source["EmployeeID"];
	        this.CurrentCategoryID = source["CurrentCategoryID"];
	        this.Status = source["Status"];
	    }
	}

}

export namespace rfidcard {
	
	export class RfidCard {
	    id: number;
	    uuid: string;
	    status: boolean;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    deactivated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new RfidCard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.uuid = source["uuid"];
	        this.status = source["status"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.deactivated_at = this.convertValues(source["deactivated_at"], null);
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

export namespace rfidcardhistory {
	
	export class RfidCardHistory {
	    id: number;
	    // Go type: time
	    assigned_at: any;
	    // Go type: time
	    unassigned_at: any;
	    employee_id: number;
	    rfid_card_id: number;
	
	    static createFrom(source: any = {}) {
	        return new RfidCardHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.assigned_at = this.convertValues(source["assigned_at"], null);
	        this.unassigned_at = this.convertValues(source["unassigned_at"], null);
	        this.employee_id = source["employee_id"];
	        this.rfid_card_id = source["rfid_card_id"];
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

