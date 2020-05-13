import { useState } from "react";

function useInputState(initValue) {
    const [state, setState] = useState(initValue);

    const handleChange = (evt) => {
        setState(evt.target.value);
    };

    const handleReset = () => {
        setState("");
    }

    return [state, handleChange, handleReset];
}

export default useInputState;