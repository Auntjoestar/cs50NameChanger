// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {gorm} from '../models';
import {models} from '../models';

export function ChangeFileNames(arg1:Array<string>,arg2:string):Promise<void>;

export function ConnectDB():Promise<gorm.DB>;

export function CreateCycle(arg1:string,arg2:string):Promise<void>;

export function CreateDB():Promise<void>;

export function CreateGroup(arg1:string,arg2:string):Promise<void>;

export function CreateProgram(arg1:string):Promise<void>;

export function CreateWeek(arg1:string,arg2:string):Promise<void>;

export function ListCycles(arg1:string):Promise<Array<string>>;

export function ListGroups(arg1:string):Promise<Array<string>>;

export function ListPrograms():Promise<Array<string>>;

export function ListWeeks(arg1:string):Promise<Array<string>>;

export function MakeNewName(arg1:string,arg2:string,arg3:string,arg4:string):Promise<string>;

export function OpenFileDialog():Promise<Array<string>>;

export function WatchCycles():Promise<Array<models.CyclesResponse>>;

export function WatchGroups():Promise<Array<models.GroupsResponse>>;

export function WatchPrograms():Promise<Array<models.ProgramsResponse>>;

export function WatchWeeks():Promise<Array<models.WeeksResponse>>;
