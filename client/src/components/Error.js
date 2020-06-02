import React from "react"
import Alert from '@material-ui/lab/Alert';

function Error(props) {
    return (
        <div className="content">
            <Alert
                className="alert"
                variant="filled"
                severity={"error"}
                onClose={() => {
                    props.setMessage("")
                }}
            >
                {props.message}
            </Alert>
        </div>
    )
}

export default Error