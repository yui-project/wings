import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { useSelector, useDispatch } from 'react-redux';
import { Telemetry, TelemetryViewIndex } from '../../../models';
import { RootState } from '../../../redux/store/RootState';
import { getTelemetryColor, getLatestTelemetries } from '../../../redux/telemetries/selectors';
import Toolbar from '@material-ui/core/Toolbar';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';
import Button from '@material-ui/core/Button';
import { setTelemetryTypePacketAction } from '../../../redux/views/actions';
import CenterFocusStrongIcon from '@material-ui/icons/CenterFocusStrong';
import IconButtonInTabs from '../../common/IconButtonInTabs';
import OpenPacketTabDialog from './OpenPacketTabDialog';

const useStyles = makeStyles((theme: Theme) => 
  createStyles({
    root: {
      padding: 10
    },
    item: {
      [theme.breakpoints.down('sm')]: {
        width: '100%'
      },
      [theme.breakpoints.up('sm')]: {
        width: '50%'
      },
      [theme.breakpoints.up('md')]: {
        width: '33.33%'
      },
    },
    tlmul: {
      paddingInlineStart: 0,
      margin: 0
    },
    tlmLayer: {
      fontSize: 'xx-small',
      display: 'block',
      color: '#ffff00',
      paddingTop: "5px"
    },
    tlmLayerOther: {
      fontSize: 'xx-small',
      display: 'block',
      paddingTop: "5px"
    },
    tlmli: {
      fontSize: 'xx-small',
      display: 'block'
    },
    tlmNormal: {
      "& span" : {
        color: theme.palette.success.main
      }
    },
    tlmColorRed: {
      "& span": {
        color: theme.palette.error.main
      }
    },
    tlmColorBlue: {
      "& span": {
        color: theme.palette.success.dark
      }
    },
    tlmColorGreen: {
      "& span": {
        color: theme.palette.success.contrastText
      }
    },
    dataTypeField: {
      fontSize: "10pt",
      textAlign:"center"
    },
    dialogPaper: {
      width: '80%',
      maxHeight: 435,
    },
    packetId: {
      color: 'white',
      fontSize: 12,
      paddingRight: 20
    },
    packetIdTitle: {
      color: '#ffff00'
    },
    tabName: {
      color: 'white',
      fontSize: 12,
      paddingRight: 20
    },
    tabNameTitle: {
      color: '#ffff00'
    }
}));

export interface PacketTabPanelProps {
  tab: TelemetryViewIndex,
  blockNum: number
}

const PacketTabPanel = (props: PacketTabPanelProps) => {
  const { tab, blockNum } = props;
  const selector = useSelector((state: RootState) => state);
  const classes = useStyles();
  const dispatch = useDispatch();
  const tlms = getLatestTelemetries(selector)[tab.name];
  const selectedTelemetries = tab.selectedTelemetries;
  const tlmClassList: string[] = [tab.name];
  const tlmColor = getTelemetryColor(selector);

  let tlmsDisplayed:Telemetry[] = [];
  tlms.forEach(tlm => {
    if (selectedTelemetries.indexOf(tlm.telemetryInfo.name) >= 0){
      tlmsDisplayed.push(tlm);
    }
  })
  const num = tlmsDisplayed.length;

  const [dataType, setDataType] = React.useState(tab.dataType);
  const [dialogOpen, setDialogOpen] = React.useState(false);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDataType((event.target as HTMLInputElement).value);
  };

  const handleOk = () => {
    dispatch(setTelemetryTypePacketAction(blockNum, dataType));
  };

  const handleDialogOpen = () => {
    setDialogOpen(true);
  }

  const handleDialogClose = () => {
    setDialogOpen(false);
  };

  const setTlmTagColor = (i: number) => {
    if (i <= 2) {
      return classes.tlmLayer;
    } else {
      return classes.tlmLayerOther;
    }
  }

  const setClassName = (dataType: string, value: string) => {
    if (dataType == "Raw") {
      return `${classes.tlmli} ${classes.tlmNormal}`;
    } else if (tlmColor.red.includes(value)) {
      return `${classes.tlmli} ${classes.tlmColorRed}`;
    } else if (tlmColor.green.includes(value)) {
      return `${classes.tlmli} ${classes.tlmColorGreen}`;
    } else if (tlmColor.blue.includes(value)) {
      return `${classes.tlmli} ${classes.tlmColorBlue}`;
    } else {
      return `${classes.tlmli} ${classes.tlmNormal}`;
    }
  }

  const showTlmData = (tlm: Telemetry) => {
    const tlmClasses = tlm.telemetryInfo.name.split('.');
    const tlmClassesDisplayed: JSX.Element[] = [];
    if (tlmClasses.length == 1){
      tlmClassesDisplayed.push(
        <li key={tlm.telemetryInfo.name} className={setClassName(tab.dataType, tlm.telemetryValue.value)}>
          {tlm.telemetryInfo.name} : <span>{(tab.dataType != "Raw")? tlm.telemetryValue.value: tlm.telemetryValue.rawValue}</span>
        </li>
      )
      tlmClassList.push(tlm.telemetryInfo.name);
    } else {
      const thisTlmClasses:string[] = [];
      tlmClasses.forEach((tlmName, i) => {
        let tlmClassesTmp = (i == 0)? tlmName :tlmClasses.slice(0,i+1).join(".");
        if (i == tlmClasses.length - 1) {
          tlmClassesDisplayed.push(
            <li key={tlm.telemetryInfo.name} className={setClassName(tab.dataType, tlm.telemetryValue.value)}>
              {<span style={{marginRight: `${10*i}px`}}></span>}
              {tlmName} : <span>{(tab.dataType != "Raw")? tlm.telemetryValue.value: tlm.telemetryValue.rawValue}</span>
            </li>
          )
        } else if (!tlmClassList.includes(tlmClassesTmp)) {
          tlmClassesDisplayed.push(
            <li key={tlm.telemetryInfo.name} className={setTlmTagColor(i)}>
              {<span style={{ marginRight: `${10 * i}px` }}></span>}
              {tlmName}
            </li>
          )
        }
        thisTlmClasses.push(tlmClassesTmp);
      })
      thisTlmClasses.forEach(thisTlmClass => { if (!tlmClassList.includes(thisTlmClass)) tlmClassList.push(thisTlmClass) });
    }
    return (
      <>
      {(tlmClassesDisplayed.map((tlmClass) => tlmClass))}
      </>
    )
  };

  return (
    <div className={classes.root}>
      <Toolbar>
        <FormControl component="fieldset">
          <Toolbar>
            <FormLabel component="legend" className={classes.dataTypeField}>Data Type</FormLabel>
            <RadioGroup aria-label="data-type" name="data-type" value={dataType} onChange={handleChange}>
              <Toolbar>
                <FormControlLabel value="Default" control={<Radio />} label="Default" />
                <FormControlLabel value="Raw" control={<Radio />} label="Raw" />
              </Toolbar>
            </RadioGroup>
          </Toolbar>
        </FormControl>
        <div className={classes.tabName}>
          <span className={classes.tabNameTitle}>Name : </span>{tab.name}
        </div>
        <div className={classes.packetId}>
          <span className={classes.packetIdTitle}>Packet Id : </span> 0x{Number(tab.packetId).toString(16)}
        </div>
        <Button onClick={handleOk} color="primary">
          SET
        </Button>
        <IconButtonInTabs onClick={handleDialogOpen}>
          <CenterFocusStrongIcon fontSize="small"/>
        </IconButtonInTabs>
      </Toolbar>
      <OpenPacketTabDialog
        blockNum={blockNum}
        classes={{ paper: classes.dialogPaper }}
        keepMounted
        open={dialogOpen}
        tab={tab}
        onClose={handleDialogClose}
      />
      <Grid
        container
        spacing={2}
      >
        <Grid item className={classes.item}>
          {tlmsDisplayed.filter((tlm, i) => i < num / 3)
          .map(tlm => (
            showTlmData(tlm)
          ))}
        </Grid>
        <Grid item className={classes.item}>
          {tlmsDisplayed.filter((tlm, i) => i >= num / 3 && i < 2 * num / 3)
          .map((tlm => (
            showTlmData(tlm)
          )))}
        </Grid>
        <Grid item className={classes.item}>
          {tlmsDisplayed.filter((tlm, i) => i >= 2 * num / 3)
          .map(tlm => (
            showTlmData(tlm)
          ))}
        </Grid>
      </Grid>
    </div>
  )
};

export default PacketTabPanel;
