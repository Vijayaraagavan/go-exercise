/*
when the function expects return value as bool and the input parameter also bool, first try to write logic so that there is only one single statement(return)

*/

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
  if knightIsAwake {        //return !knightIsAwake (simple return statement)
        return false
    } else {
    return true
    }
	panic("Please implement the CanFastAttack() function")
}
//-----------------------------------------------------------------------

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    if (knightIsAwake || archerIsAwake || prisonerIsAwake) {
        return true
    } else {
    return false
    }
	panic("Please implement the CanSpy() function")
}
// in question, note "atleast one of the ", so 
//return knightIsAwake || archerIsAwake || prisonerIsAwake;
//---------------------------------------------------------------------------------

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if (prisonerIsAwake == true && archerIsAwake != true) {
        return true
    } else {
    return false
    }
	panic("Please implement the CanSignalPrisoner() function")
}
//return prisonerIsAwake && !archerIsAwake;
// this is deMorgan's theorem ( boolean logic)
//----------------------------------------------------------------------------------


//CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    if (knightIsAwake!=true && archerIsAwake!=true && prisonerIsAwake==true) || (petDogIsPresent == true && archerIsAwake == false) {
        return true
    } else {
    return false
    }
	panic("Please implement the CanFreePrisoner() function")
}
//return !archerIsAwake && ((prisonerIsAwake && !knightIsAwake) || petDogIsPresent);
//here in both case archer should be in sleep

----------------------------------------------------------------------------
