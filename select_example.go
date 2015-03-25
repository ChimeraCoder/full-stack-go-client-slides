package main


// START OMIT
select {
    case <-algorithm1:
        // algorithm1 was faster
    case <-algorithm2:
        // algorithm2 was faster
}
// END OMIT
