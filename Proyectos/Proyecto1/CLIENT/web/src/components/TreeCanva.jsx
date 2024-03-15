
import { useEffect, useRef } from 'react';
import * as d3 from 'd3';
import 'd3-graphviz';

const TreeCanva = ({data}) => {

  const graphRef = useRef();

  useEffect(() => {
    if (data) {
      const dot = generateDot(data[0]);
      const fdot = finalDot(data);
      const graph = d3.select(graphRef.current).graphviz()
      graph
        .renderDot(fdot)
        .on('end', () => {
          // Adjust SVG size to match container size
          const svg = d3.select(graphRef.current).select('svg');
          svg.attr('width', '100%').attr('height', '100%');
        });;
    }
  }, [data]);


    

  const generateDot = (node) => {
    let dot = `
      ${node.pid} [label="${node.name}"];
    `;
    if (node.child) {
      node.child.forEach((child) => {
        dot += `${node.pid} -> ${child.pid};\n`;
        dot += generateDot(child);
      });
    }

    return dot;
  };

  const finalDot =(data)=> {
    return `
      digraph {
        node [style=filled, shape=box, fontname=Arial, fontsize=12, fontcolor=white, color=black, fillcolor=black, width=0.05, height=0.05, margin=0.05, penwidth=2];
        ${generateDot(data[0])}
      }
    `;

  } 

  return(
    
    
    <div ref={graphRef} className="w-full h-full">
    </div>
      
  );
};

export default TreeCanva